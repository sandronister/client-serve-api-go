package priceservice

import "context"

func (s *service) GetQuote(ctx context.Context) {
	currentQuote, err := s.usecase.GetQuote(ctx)
	if err != nil {
		panic(err)
	}

	err = s.repository.Insert("cotacao.txt", currentQuote)
	if err != nil {
		panic(err)
	}
}
