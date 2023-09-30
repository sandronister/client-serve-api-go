package main

import (
	"net/http"

	"github.com/gofor-little/env"
	"github.com/gorilla/mux"
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
	conn := sqllite.GetConnection()
	priceUseCase := di.ConfigPriceDI(conn)
	defer conn.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", priceUseCase.Get).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
