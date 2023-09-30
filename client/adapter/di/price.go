package di

import (
	"github.com/sandronister/client-serve-api-go/client/adapter/priceservice"
	"github.com/sandronister/client-serve-api-go/client/core/domain"
	"github.com/sandronister/client-serve-api-go/client/core/domain/usecase/priceusecase"
	"github.com/sandronister/client-serve-api-go/client/file/pricerepository"
)

func ConfigPriceDI() domain.PriceService {
	repository := pricerepository.New()
	usecase := priceusecase.New()
	service := priceservice.New(usecase, repository)
	return service
}
