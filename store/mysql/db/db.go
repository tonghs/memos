package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/usememos/memos/server/profile"
	"github.com/usememos/memos/store"
	"strings"
)

//go:embed migration
var migrationFS embed.FS

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

	switch db.profile.DBDriver {
	case store.MySQLDriver, store.PostgresDriver:
		db.DBInstance, err = sql.Open(store.MySQLDriver, db.profile.DSN)
	default:
		return fmt.Errorf("failed to open db, err: unknown db driver: %s", db.profile.DBDriver)
	}

	if err != nil {
		return fmt.Errorf("failed to open db with dsn: %s, err: %w", db.profile.DSN, err)
	}

	if err = db.DBInstance.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to connect to db instance with dsn: %s, err: %w", db.profile.DSN, err)
	}

	// if database table not exists, we should migrate the database.
	if _, err := db.DBInstance.ExecContext(ctx, "desc table user"); err != nil {
		if err, ok := err.(*mysql.MySQLError); ok {
			if int(err.Number) == 1146 {
				if err := db.applyLatestSchema(ctx); err != nil {
					return fmt.Errorf("failed to apply latest schema: %w", err)
				}
			}
		}
	}

	// TODO migration and seed

	return nil
}

const (
	latestSchemaFileName = "LATEST__SCHEMA.sql"
)

func (db *DB) applyLatestSchema(ctx context.Context) error {
	schemaMode := "dev"
	if db.profile.Mode == "prod" {
		schemaMode = "prod"
	}
	latestSchemaPath := fmt.Sprintf("%s/%s/%s", "migration", schemaMode, latestSchemaFileName)
	buf, err := migrationFS.ReadFile(latestSchemaPath)
	if err != nil {
		return fmt.Errorf("failed to read latest schema %q, error %w", latestSchemaPath, err)
	}
	for _, stmt := range strings.Split(string(buf), ";") {
		if strings.TrimSpace(stmt) == "" {
			continue
		}

		if _, err := db.DBInstance.ExecContext(ctx, stmt); err != nil {
			return fmt.Errorf("migrate error: statement:%s err=%w", stmt, err)
		}
	}
	return nil
}
