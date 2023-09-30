package priceservice

import "github.com/sandronister/client-serve-api-go/client/core/domain"

type service struct {
	usecase    domain.PriceUseCase
	repository domain.PriceRepository
}

func New(usecase domain.PriceUseCase, repository domain.PriceRepository) domain.PriceService {
	return &service{
		usecase:    usecase,
		repository: repository,
	}
}
