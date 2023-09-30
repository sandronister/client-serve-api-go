package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sandronister/client-serve-api-go/server/adapter/http/di"
	"github.com/sandronister/client-serve-api-go/server/adapter/sqllite"
)

func main() {
	conn := sqllite.GetConnection()
	priceUseCase := di.ConfigPriceDI(conn)
	defer conn.Close()
	router := mux.NewRouter()
	router.HandleFunc("/cotacao", priceUseCase.Get).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
