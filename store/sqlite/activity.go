package sqlite

import (
	"context"
	"database/sql"
	"github.com/usememos/memos/api"
	"github.com/usememos/memos/store"
)

// CreateActivity creates an instance of Activity.
func (s *Store) CreateActivity(ctx context.Context, create *api.ActivityCreate) (*api.Activity, error) {
	if s.profile.Mode == "prod" {
		return nil, nil
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, store.FormatError(err)
	}
	defer tx.Rollback()

	activityRaw, err := createActivity(ctx, tx, create)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, store.FormatError(err)
	}

	activity := activityRaw.ToActivity()
	return activity, nil
}

// createActivity creates a new activity.
func createActivity(ctx context.Context, tx *sql.Tx, create *api.ActivityCreate) (*store.ActivityRaw, error) {
	query := `
		INSERT INTO activity (
			creator_id, 
			type, 
			level, 
			payload
		)
		VALUES (?, ?, ?, ?)
		RETURNING id, type, level, payload, creator_id, created_ts
	`
	var activityRaw store.ActivityRaw
	if err := tx.QueryRowContext(ctx, query, create.CreatorID, create.Type, create.Level, create.Payload).Scan(
		&activityRaw.ID,
		&activityRaw.Type,
		&activityRaw.Level,
		&activityRaw.Payload,
		&activityRaw.CreatorID,
		&activityRaw.CreatedTs,
	); err != nil {
		return nil, store.FormatError(err)
	}

	return &activityRaw, nil
}
