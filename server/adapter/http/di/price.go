package di

import (
	"sandroni.fullcycle.server/adapter/http/priceservice"
	"sandroni.fullcycle.server/adapter/sqllite"
	"sandroni.fullcycle.server/adapter/sqllite/pricerepository"
	"sandroni.fullcycle.server/core/domain"
	"sandroni.fullcycle.server/core/domain/usecase/priceusecase"
)

func ConfigPriceDI(conn sqllite.IDatabase) domain.PriceService {
	priceRepository := pricerepository.New(conn)
	priceUseCase := priceusecase.New(priceRepository)
	priceService := priceservice.New(priceUseCase)
	return priceService
}
