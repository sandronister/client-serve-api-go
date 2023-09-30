package pricerepository

import (
	"github.com/sandronister/client-serve-api-go/server/adapter/sqllite"
	"github.com/sandronister/client-serve-api-go/server/core/domain"
)

type repository struct {
	db sqllite.IDatabase
}

func initTable(db sqllite.IDatabase) error {
	createTable := `CREATE TABLE price(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"code" TEXT,
		"code_in" TEXT,
		"name" TEXT,
		"high" TEXT,
		"low" TEXT,
		"var_bid" TEXT,
		"pct_change" TEXT,
		"bid" TEXT,
		"ask" TEXT,
		"timestamp" TEXT,
		"create_date" TEXT
	)`

	stmt, err := db.Prepare(createTable)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}
	return nil
}

func New(db sqllite.IDatabase) domain.PriceRepository {
	initTable(db)
	return &repository{
		db: db,
	}
}
