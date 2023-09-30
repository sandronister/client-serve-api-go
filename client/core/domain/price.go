package domain

import "context"

type Price struct {
	Bid  string
	Data string
}

type PriceUseCase interface {
	GetQuote(ctx context.Context) ([]byte, error)
}

type PriceRepository interface {
	Insert(filename string, content []byte) error
}

type PriceService interface {
	GetQuote(ctx context.Context)
}
