package sqllite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofor-little/env"
	_ "github.com/mattn/go-sqlite3"
)

type IDatabase interface {
	Close() error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
}

func GetConnection() *sql.DB {
	DSN := env.Get("DSN", "FAIL")

	if DSN == "FAIL" {
		panic(errors.New("DSN not defined"))
	}

	db, err := sql.Open("sqlite3", DSN)

	if err != nil {
		panic(err)
	}

	return db
}
