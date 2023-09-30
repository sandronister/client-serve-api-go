package main

import (
	"context"
	"time"

	"github.com/sandronister/client-serve-api-go/client/adapter/di"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	priceservice := di.ConfigPriceDI()

	priceservice.GetQuote(ctx)
}
