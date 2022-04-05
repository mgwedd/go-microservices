package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mgwedd/go-microservices/product-api/data"
)

type Products struct {
	logger *log.Logger
}

type ProductKey struct{}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJson(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}
}

func (product *Products) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle POST product")

	prod := request.Context().Value(ProductKey{}).(data.Product)

	data.AddProduct(&prod)
	product.logger.Printf("Product: %#v", prod)
}

func (product *Products) UpdateProduct(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle PUT product")

	reqParams := mux.Vars(request)
	id, err := strconv.Atoi(reqParams["id"])
	if err != nil {
		http.Error(responseWriter, "Unable to parse ID", http.StatusBadRequest)
	}

	prod := request.Context().Value(ProductKey{}).(data.Product)

	updateErr := data.UpdateProduct(id, &prod)
	if updateErr == data.ErrProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if updateErr != nil {
		http.Error(responseWriter, "Cannot update product", http.StatusBadRequest)
		return
	}

	product.logger.Printf("Product updated: %#v", prod)
}

func (product *Products) DeleteProduct(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle DELETE product")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(responseWriter, "Unable to parse ID", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}

	updateErr := data.DeleteProduct(id, prod)
	if updateErr == data.ErrProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if updateErr != nil {
		http.Error(responseWriter, "Cannot delete product", http.StatusBadRequest)
		return
	}

	product.logger.Printf("Product deleted: %#v", prod)
}

func (products Products) ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {

		products.logger.Println("[DEBUG] validating request")
		prod := data.Product{}

		parseErr := prod.FromJSON(request.Body)
		if parseErr != nil {
			products.logger.Println("[ERROR] deserializing product during request validation", parseErr)
			http.Error(responseWriter, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		// Validate the product
		err := prod.Validate()
		if err != nil {
			products.logger.Println("[ERROR] Product request validation failed", parseErr)
			http.Error(responseWriter, fmt.Sprintf("Product validation failed: %s", err.Error()), http.StatusBadRequest)
			return
		}

		// Add the product to the request context
		ctx := context.WithValue(request.Context(), ProductKey{}, prod)
		reqCtx := request.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(responseWriter, reqCtx)
	})
}
