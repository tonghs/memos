package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/usememos/memos/store"
	"strings"

	"github.com/usememos/memos/api"
	"github.com/usememos/memos/common"
)

func (s *Store) UpsertSystemSetting(ctx context.Context, upsert *api.SystemSettingUpsert) (*api.SystemSetting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	systemSettingRaw, err := upsertSystemSetting(ctx, tx, upsert)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	systemSetting := systemSettingRaw.ToSystemSetting()

	return systemSetting, nil
}

func (s *Store) FindSystemSettingList(ctx context.Context, find *api.SystemSettingFind) ([]*api.SystemSetting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	systemSettingRawList, err := findSystemSettingList(ctx, tx, find)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	list := []*api.SystemSetting{}
	for _, raw := range systemSettingRawList {
		list = append(list, raw.ToSystemSetting())
	}

	return list, nil
}

func (s *Store) FindSystemSetting(ctx context.Context, find *api.SystemSettingFind) (*api.SystemSetting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	systemSettingRawList, err := findSystemSettingList(ctx, tx, find)
	if err != nil {
		return nil, err
	}

	if len(systemSettingRawList) == 0 {
		return nil, &common.Error{Code: common.NotFound, Err: fmt.Errorf("not found")}
	}

	return systemSettingRawList[0].ToSystemSetting(), nil
}

func upsertSystemSetting(ctx context.Context, tx *sql.Tx, upsert *api.SystemSettingUpsert) (*store.SystemSettingRaw, error) {
	query := `
		INSERT INTO system_setting (
			name, value, description
		)
		VALUES (?, ?, ?)
		ON CONFLICT(name) DO UPDATE 
		SET
			value = EXCLUDED.value,
			description = EXCLUDED.description
		RETURNING name, value, description
	`
	var systemSettingRaw store.SystemSettingRaw
	if err := tx.QueryRowContext(ctx, query, upsert.Name, upsert.Value, upsert.Description).Scan(
		&systemSettingRaw.Name,
		&systemSettingRaw.Value,
		&systemSettingRaw.Description,
	); err != nil {
		return nil, store.FormatError(err)
	}

	return &systemSettingRaw, nil
}

func findSystemSettingList(ctx context.Context, tx *sql.Tx, find *api.SystemSettingFind) ([]*store.SystemSettingRaw, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if find.Name.String() != "" {
		where, args = append(where, "name = ?"), append(args, find.Name.String())
	}

	query := `
		SELECT
			name,
		  value,
			description
		FROM system_setting
		WHERE ` + strings.Join(where, " AND ")
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer rows.Close()

	systemSettingRawList := make([]*store.SystemSettingRaw, 0)
	for rows.Next() {
		var systemSettingRaw store.SystemSettingRaw
		if err := rows.Scan(
			&systemSettingRaw.Name,
			&systemSettingRaw.Value,
			&systemSettingRaw.Description,
		); err != nil {
			return nil, store.FormatError(err)
		}

		systemSettingRawList = append(systemSettingRawList, &systemSettingRaw)
	}

	if err := rows.Err(); err != nil {
		return nil, store.FormatError(err)
	}

	return systemSettingRawList, nil
}
