package priceusecase

import "sandroni.fullcycle.server/core/domain"

type usecase struct {
	repository domain.PriceRepository
}

func New(repository domain.PriceRepository) domain.PriceUseCase {
	return &usecase{
		repository: repository,
	}
}
