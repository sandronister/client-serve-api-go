package priceusecase

import "github.com/sandronister/client-serve-api-go/server/core/domain"

type usecase struct {
	repository domain.PriceRepository
}

func New(repository domain.PriceRepository) domain.PriceUseCase {
	return &usecase{
		repository: repository,
	}
}
