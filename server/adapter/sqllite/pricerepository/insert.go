package pricerepository

import (
	"context"

	"github.com/sandronister/client-serve-api-go/server/core/domain"
)

func (r *repository) Insert(ctx context.Context, price *domain.Price) error {
	sqlCommand := "Insert into price(code,code_in,name,high,low,var_bid,pct_change,bid,ask,timestamp,create_date) values (?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.PrepareContext(ctx, sqlCommand)

	if err != nil {
		return err
	}

	defer stmt.Close()

	values := &price.Usdbrl
	_, err = stmt.ExecContext(ctx, values.Code, values.Codein, values.Name, values.High, values.Low, values.VarBid, values.PctChange, values.Bid, values.Ask, values.Timestamp, values.CreateDate)

	if err != nil {
		return err
	}

	return nil
}
