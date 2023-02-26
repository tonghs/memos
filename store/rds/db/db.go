package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/usememos/memos/server/profile"
	"github.com/usememos/memos/store"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	// sqlite db connection instance
	DBInstance *sql.DB
	profile    *profile.Profile
}

// Instance returns the db instance.
func (db *DB) Instance() *sql.DB {
	return db.DBInstance
}

// NewDB returns a new instance of DB associated with the given datasource name.
func NewDB(profile *profile.Profile) *DB {
	db := &DB{
		profile: profile,
	}
	return db
}

func (db *DB) Open(ctx context.Context) (err error) {
	if db.profile.DSN == "" {
		return fmt.Errorf("dsn required")
	}

	var _db *gorm.DB
	switch db.profile.DBDriver {
	case store.MySQLDriver:
		_db, err = gorm.Open(mysql.Open(db.profile.DSN), &gorm.Config{})
	case store.PostgresDriver:
		_db, err = gorm.Open(postgres.Open(db.profile.DSN), &gorm.Config{})
	default:
		return fmt.Errorf("failed to open db, err: unknown db driver: %s", db.profile.DBDriver)
	}

	if err != nil {
		return fmt.Errorf("failed to open db with dsn: %s, err: %w", db.profile.DSN, err)
	}

	db.DBInstance, err = _db.DB()
	if err != nil {
		return fmt.Errorf("failed to get db instance with dsn: %s, err: %w", db.profile.DSN, err)
	}

	if err = db.DBInstance.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to connect to db instance with dsn: %s, err: %w", db.profile.DSN, err)
	}

	return nil
}
