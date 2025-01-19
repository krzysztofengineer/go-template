package database

import (
	"database/sql"
	"go-template/schema"

	"github.com/pressly/goose/v3"

	_ "modernc.org/sqlite"
)

func New(dsn string) *sql.DB {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(schema.FS)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}

	return db
}
