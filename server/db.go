package server

import (
	"fmt"

	"github.com/usememos/memos/server/profile"
	"github.com/usememos/memos/store"
	rdsDB "github.com/usememos/memos/store/rds/db"
	sqliteDB "github.com/usememos/memos/store/sqlite/db"
)

func NewDB(profile *profile.Profile) store.DB {
	switch profile.DBDriver {
	case store.SQLiteDriver:
		return sqliteDB.NewDB(profile)
	case store.MySQLDriver, store.PostgresDriver:
		return rdsDB.NewDB(profile)
	default:
		panic(fmt.Sprintf("unknown db driver: %s", profile.DBDriver))
	}
}
