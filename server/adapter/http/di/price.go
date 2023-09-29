package di

import (
	"sandroni.fullcycle.server/adapter/sqllite"
	"sandroni.fullcycle.server/adapter/sqllite/pricerepository"
	"sandroni.fullcycle.server/core/domain"
	"sandroni.fullcycle.server/core/domain/usecase/priceusecase"
)

func ConfigPriceDI(conn sqllite.IDatabase) domain.PriceUseCase {
	priceRepository := pricerepository.New(conn)
	priceUseCase := priceusecase.New(priceRepository)
	return priceUseCase
}
