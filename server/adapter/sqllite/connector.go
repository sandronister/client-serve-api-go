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
}

func GetConnection() (*sql.DB, error) {
	DSN := env.Get("DSN", "FAIL")

	if DSN == "FAIL" {
		return nil, errors.New("DSN not defined")
	}

	db, err := sql.Open("sqllite3", "./price.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}
