package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (product *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

type Products []*Product

func (products *Products) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(products)
}

func GetProducts() Products {
	return productList
}

func AddProduct(product *Product) {
	product.ID = GetNextID()
	productList = append(productList, product)
}

func UpdateProduct(id int, product *Product) error {
	_, posIdx, err := FindProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[posIdx] = product
	return nil
}

func FindProduct(id int) (*Product, int, error) {
	for posIdx, product := range productList {
		if product.ID == id {
			return product, posIdx, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func GetNextID() int {
	prodList := productList[len(productList)-1]
	return prodList.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "xyz321",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
