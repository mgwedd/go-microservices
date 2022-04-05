package data

import "testing"

func TestChecksValidation(t *testing.T) {
	prod := &Product{
		Name:  "Test Coffee",
		Price: 1.50,
		SKU:   "abc-def-higj",
	}

	err := prod.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
