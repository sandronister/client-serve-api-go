package di

import (
	"github.com/sandronister/client-serve-api-go/server/adapter/http/priceservice"
	"github.com/sandronister/client-serve-api-go/server/adapter/sqllite"
	"github.com/sandronister/client-serve-api-go/server/adapter/sqllite/pricerepository"
	"github.com/sandronister/client-serve-api-go/server/core/domain"
	"github.com/sandronister/client-serve-api-go/server/core/domain/usecase/priceusecase"
)

func ConfigPriceDI(conn sqllite.IDatabase) domain.PriceService {
	priceRepository := pricerepository.New(conn)
	priceUseCase := priceusecase.New(priceRepository)
	priceService := priceservice.New(priceUseCase)
	return priceService
}
