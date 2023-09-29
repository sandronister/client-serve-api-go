package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type USResult struct {
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

type PriceResult struct {
	Usdbrl USResult `json:"USDBRL"`
}

func getPrice(ctx context.Context) PriceResult {
	price := PriceResult{}

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &price)

	if err != nil {
		panic(err)
	}

	return price

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	result := getPrice(ctx)
	json.NewEncoder(os.Stdout).Encode(result)
}
