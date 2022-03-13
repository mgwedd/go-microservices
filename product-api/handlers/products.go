package handlers

import (
	"log"
	"net/http"

	"github.com/mgwedd/go-microservices/product-api/data"
)

type Product struct {
	logger *log.Logger
}

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Product {
	return &Product{logger}
}

func (products *Product) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJson(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJson(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
