package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/usememos/memos/common"
	"github.com/usememos/memos/store"
)

func (s *Store) CreateIdentityProvider(ctx context.Context, create *store.IdentityProviderMessage) (*store.IdentityProviderMessage, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	var configBytes []byte
	if create.Type == store.IdentityProviderOAuth2 {
		configBytes, err = json.Marshal(create.Config.OAuth2Config)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unsupported idp type %s", string(create.Type))
	}
	query := `
		INSERT INTO idp (
			name,
			type,
			identifier_filter,
			config
		)
		VALUES (?, ?, ?, ?)
		RETURNING id
	`
	if err := tx.QueryRowContext(
		ctx,
		query,
		create.Name,
		create.Type,
		create.IdentifierFilter,
		string(configBytes),
	).Scan(
		&create.ID,
	); err != nil {
		return nil, store.FormatError(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, store.FormatError(err)
	}
	identityProviderMessage := create
	s.idpCache.Store(identityProviderMessage.ID, identityProviderMessage)
	return identityProviderMessage, nil
}

func (s *Store) ListIdentityProviders(ctx context.Context, find *store.FindIdentityProviderMessage) ([]*store.IdentityProviderMessage, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	list, err := listIdentityProviders(ctx, tx, find)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		s.idpCache.Store(item.ID, item)
	}
	return list, nil
}

func (s *Store) GetIdentityProvider(ctx context.Context, find *store.FindIdentityProviderMessage) (*store.IdentityProviderMessage, error) {
	if find.ID != nil {
		if cache, ok := s.idpCache.Load(*find.ID); ok {
			return cache.(*store.IdentityProviderMessage), nil
		}
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	list, err := listIdentityProviders(ctx, tx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, &common.Error{Code: common.NotFound, Err: fmt.Errorf("not found")}
	}

	identityProviderMessage := list[0]
	s.idpCache.Store(identityProviderMessage.ID, identityProviderMessage)
	return identityProviderMessage, nil
}

func (s *Store) UpdateIdentityProvider(ctx context.Context, update *store.UpdateIdentityProviderMessage) (*store.IdentityProviderMessage, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	set, args := []string{}, []interface{}{}
	if v := update.Name; v != nil {
		set, args = append(set, "name = ?"), append(args, *v)
	}
	if v := update.IdentifierFilter; v != nil {
		set, args = append(set, "identifier_filter = ?"), append(args, *v)
	}
	if v := update.Config; v != nil {
		var configBytes []byte
		if update.Type == store.IdentityProviderOAuth2 {
			configBytes, err = json.Marshal(update.Config.OAuth2Config)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unsupported idp type %s", string(update.Type))
		}
		set, args = append(set, "config = ?"), append(args, string(configBytes))
	}
	args = append(args, update.ID)

	query := `
		UPDATE idp
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING id, name, type, identifier_filter, config
	`
	var identityProviderMessage store.IdentityProviderMessage
	var identityProviderConfig string
	if err := tx.QueryRowContext(ctx, query, args...).Scan(
		&identityProviderMessage.ID,
		&identityProviderMessage.Name,
		&identityProviderMessage.Type,
		&identityProviderMessage.IdentifierFilter,
		&identityProviderConfig,
	); err != nil {
		return nil, store.FormatError(err)
	}
	if identityProviderMessage.Type == store.IdentityProviderOAuth2 {
		oauth2Config := &store.IdentityProviderOAuth2Config{}
		if err := json.Unmarshal([]byte(identityProviderConfig), oauth2Config); err != nil {
			return nil, err
		}
		identityProviderMessage.Config = &store.IdentityProviderConfig{
			OAuth2Config: oauth2Config,
		}
	} else {
		return nil, fmt.Errorf("unsupported idp type %s", string(identityProviderMessage.Type))
	}
	if err := tx.Commit(); err != nil {
		return nil, store.FormatError(err)
	}
	s.idpCache.Store(identityProviderMessage.ID, identityProviderMessage)
	return &identityProviderMessage, nil
}

func (s *Store) DeleteIdentityProvider(ctx context.Context, delete *store.DeleteIdentityProviderMessage) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return store.FormatError(err)
	}
	defer tx.Rollback()

	where, args := []string{"id = ?"}, []interface{}{delete.ID}
	stmt := `DELETE FROM idp WHERE ` + strings.Join(where, " AND ")
	result, err := tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return store.FormatError(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return &common.Error{Code: common.NotFound, Err: fmt.Errorf("idp not found")}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	s.idpCache.Delete(delete.ID)
	return nil
}

func listIdentityProviders(ctx context.Context, tx *sql.Tx, find *store.FindIdentityProviderMessage) ([]*store.IdentityProviderMessage, error) {
	where, args := []string{"TRUE"}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, fmt.Sprintf("id = $%d", len(args)+1)), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, `
		SELECT
			id,
			name,
			type,
			identifier_filter,
			config
		FROM idp
		WHERE `+strings.Join(where, " AND ")+` ORDER BY id ASC`,
		args...,
	)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer rows.Close()

	var identityProviderMessages []*store.IdentityProviderMessage
	for rows.Next() {
		var identityProviderMessage store.IdentityProviderMessage
		var identityProviderConfig string
		if err := rows.Scan(
			&identityProviderMessage.ID,
			&identityProviderMessage.Name,
			&identityProviderMessage.Type,
			&identityProviderMessage.IdentifierFilter,
			&identityProviderConfig,
		); err != nil {
			return nil, store.FormatError(err)
		}
		if identityProviderMessage.Type == store.IdentityProviderOAuth2 {
			oauth2Config := &store.IdentityProviderOAuth2Config{}
			if err := json.Unmarshal([]byte(identityProviderConfig), oauth2Config); err != nil {
				return nil, err
			}
			identityProviderMessage.Config = &store.IdentityProviderConfig{
				OAuth2Config: oauth2Config,
			}
		} else {
			return nil, fmt.Errorf("unsupported idp type %s", string(identityProviderMessage.Type))
		}
		identityProviderMessages = append(identityProviderMessages, &identityProviderMessage)
	}

	return identityProviderMessages, nil
}
