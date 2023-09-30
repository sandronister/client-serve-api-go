package priceservice

import "github.com/sandronister/client-serve-api-go/server/core/domain"

type service struct {
	usecase domain.PriceUseCase
}

func New(usecase domain.PriceUseCase) domain.PriceService {
	return &service{
		usecase: usecase,
	}
}
