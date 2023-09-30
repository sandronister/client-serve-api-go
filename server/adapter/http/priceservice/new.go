package priceservice

import "sandroni.fullcycle.server/core/domain"

type service struct {
	usecase domain.PriceUseCase
}

func New(usecase domain.PriceUseCase) domain.PriceService {
	return &service{
		usecase: usecase,
	}
}
