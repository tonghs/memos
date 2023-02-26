package server

import (
	"database/sql"
	"fmt"

	"github.com/usememos/memos/server/profile"
	"github.com/usememos/memos/store"
	"github.com/usememos/memos/store/sqlite"
)

// NewStore creates a new instance of Store.
func NewStore(db *sql.DB, profile *profile.Profile) store.Store {
	switch profile.DBDriver {
	case store.SQLiteDriver:
		return sqlite.New(db, profile)
	// case store.MySQLDriver, store.PostgresDriver:
	//	return rds.New(db, profile)
	default:
		panic(fmt.Sprintf("unknown db driver: %s", profile.DBDriver))
	}
}
