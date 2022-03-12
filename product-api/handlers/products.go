package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mgwedd/go-microservices/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	productsList := data.GetProducts()
	data, err := json.Marshal(productsList)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON", http.StatusInternalServerError)
	}
	responseWriter.Write(data)
}
