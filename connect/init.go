package connect

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun/extra/bundebug"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"templater/config"
)

func InitDB(conf *config.Config) *bun.DB {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", conf.DbDb, conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose()))
	_, _ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	return db
}
