package product_test

import (
	"testing"

	. "github.com/joaovictornsv/shopping-cart-go/product"
)

func TestProductName(t *testing.T) {
	name := "Arroz"
	productToTest := NewProduct(name, 3.99)

	productToTestName := productToTest.GetName()
	if productToTestName != name {
		t.Errorf("Name not match!\nExpected: \"%s\"\nReceive: \"%s\"",
			name,
			productToTestName,
		)
	}
}

func TestProductPrice(t *testing.T) {
	price := 3.99
	productToTest := NewProduct("Arroz", price)

	productToTestPrice := productToTest.GetPrice()
	if productToTestPrice != price {
		t.Errorf("Price not match!\nExpected: \"%f\"\nReceive: \"%f\"",
			price,
			productToTestPrice,
		)
	}
}
