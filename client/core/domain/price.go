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
	saveFile(filename string, content []byte) error
	writeFile(filename string, body []byte) error
	appendFile(filename string, body []byte) error
	fileExists(filename string) bool
}

type PriceService interface {
	GetQuote(ctx context.Context) ([]byte, error)
}
