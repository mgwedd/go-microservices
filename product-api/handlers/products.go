package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if request.Method == http.MethodPost {
		products.addProduct(responseWriter, request)
		return
	}

	if request.Method == http.MethodPut {
		id, err := ParseIDParam(request)
		if err != nil {
			http.Error(responseWriter, "invalid product ID", http.StatusBadRequest)
		}
		products.updateProduct(id, responseWriter, request)
		return
	}

	if request.Method == http.MethodDelete {
		id, err := ParseIDParam(request)
		if err != nil {
			http.Error(responseWriter, "invalid product ID", http.StatusBadRequest)
		}
		products.deleteProduct(id, responseWriter, request)
		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

var ErrBadParams = fmt.Errorf("bad request parameters")

func ParseIDParam(request *http.Request) (int, error) {
	regex := regexp.MustCompile(`/([0-9]+)`)
	idMatchGroup := regex.FindAllStringSubmatch(request.URL.Path, -1)

	if len(idMatchGroup) != 1 {
		return -1, ErrBadParams
	}

	if len(idMatchGroup[0]) != 2 {
		return -1, ErrBadParams
	}

	idStr := idMatchGroup[0][1]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return -1, ErrBadParams
	}

	return id, nil

}

func (products *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJson(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (product *Products) addProduct(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle POST product")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)
	product.logger.Printf("Product: %#v", prod)
}

func (product *Products) updateProduct(id int, responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle PUT product")

	prod := &data.Product{}

	parseErr := prod.FromJSON(request.Body)
	if parseErr != nil {
		http.Error(responseWriter, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	updateErr := data.UpdateProduct(id, prod)
	if updateErr == data.ErrProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
	}

	if updateErr != nil {
		http.Error(responseWriter, "Cannot update product", http.StatusBadRequest)
	}

	product.logger.Printf("Product updated: %#v", prod)
}

func (product *Products) deleteProduct(id int, responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle DELETE product")

	prod := &data.Product{}

	updateErr := data.DeleteProduct(id, prod)
	if updateErr == data.ErrProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
	}

	if updateErr != nil {
		http.Error(responseWriter, "Cannot delete product", http.StatusBadRequest)
	}

	product.logger.Printf("Product deleted: %#v", prod)
}
