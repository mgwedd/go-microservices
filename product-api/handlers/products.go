package handlers

import (
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
	if request.Method == http.MethodGet {
		products.getProducts(responseWriter, request)
		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJson(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
