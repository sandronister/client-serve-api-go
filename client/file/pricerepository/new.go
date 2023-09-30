package pricerepository

import "github.com/sandronister/client-serve-api-go/client/core/domain"

type repository struct{}

func New() domain.PriceRepository {
	return &repository{}
}
