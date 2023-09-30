package priceservice

import (
	"encoding/json"
	"log"
	"net/http"
)

func (service *service) sendError(response http.ResponseWriter, err error) {
	log.Println(err)
	response.WriteHeader(500)
	response.Write([]byte("Ocorreu um erro"))
}

func (service *service) Get(response http.ResponseWriter, request *http.Request) {
	currentQuote, err := service.usecase.Get()
	if err != nil {
		service.sendError(response, err)
		return
	}
	err = service.usecase.Insert(*currentQuote)
	if err != nil {
		service.sendError(response, err)
		return
	}

	json.NewEncoder(response).Encode(currentQuote)
}
