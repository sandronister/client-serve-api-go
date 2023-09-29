package main

import (
	"fmt"

	"github.com/gofor-little/env"
	"sandroni.fullcycle.server/adapter/http/di"
	"sandroni.fullcycle.server/adapter/sqllite"
)

func init() {
	err := env.Load("../../.env")
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := sqllite.GetConnection()
	if err != nil {
		panic(err)
	}
	priceUseCase := di.ConfigPriceDI(conn)

	result, err := priceUseCase.Get()
	if err != nil {
		panic(err)
	}
	err = priceUseCase.Insert(*result)

	if err != nil {
		panic(err)
	}

	fmt.Println("Salvo com sucesso")
}
