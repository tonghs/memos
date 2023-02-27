package server

import (
	"database/sql"
	"fmt"
	"github.com/usememos/memos/store/rds"

	"github.com/usememos/memos/server/profile"
	"github.com/usememos/memos/store"
	"github.com/usememos/memos/store/sqlite"
)

// newStore creates a new instance of Store.
func newStore(db *sql.DB, profile *profile.Profile) store.Store {
	switch profile.DBDriver {
	case store.SQLiteDriver:
		return sqlite.New(db, profile)
	case store.MySQLDriver, store.PostgresDriver:
		return rds.New(db, profile)
	default:
		panic(fmt.Sprintf("unknown db driver: %s", profile.DBDriver))
	}
}
