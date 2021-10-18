package sdk

import (
	"testing"
)

func TestGetProducts(t *testing.T) {
	product := GetProducts()

	if len(product) != 1 {
		t.Fail()
	}
	if product[0].Name != "Product A" {
		t.Fail()
	}
}
