package domain

import "context"

type Price struct {
	Bid  string
	Data string
}

type PriceUseCase interface {
	GetQuote(ctx context.Context) (string, error)
}

type PriceRepository interface {
	Insert(filename, content string) error
}

type PriceService interface {
	GetQuote(ctx context.Context)
}
