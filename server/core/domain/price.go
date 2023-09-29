package domain

import "net/http"

type usResult struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Price struct {
	Usdbrl usResult `json:"USDBRL"`
}

type PriceUseCase interface {
	Get() (Price, error)
	Save(price Price) error
}

type PriceRepository interface {
	Save(price Price) error
}

type PriceService interface {
	Get(response http.ResponseWriter, request *http.Request)
}
