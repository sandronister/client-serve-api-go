package priceusecase

import (
	"context"
	"time"

	"sandroni.fullcycle.server/core/domain"
)

func (u *usecase) Insert(price domain.Price) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err := u.repository.Insert(ctx, price)

	if err != nil {
		return err
	}

	return nil
}
