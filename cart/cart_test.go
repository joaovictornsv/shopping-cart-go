package cart_test

import (
	"testing"

	. "github.com/joaovictornsv/shopping-cart-go/cart"
	"github.com/joaovictornsv/shopping-cart-go/product"
)

func TestCartProductsQuantity(t *testing.T) {
	cartToTest := NewCart()
	productToBeAdd := product.NewProduct("Arroz", 3.99)

	cartToTest.AddProduct(productToBeAdd)

	productQuantity := cartToTest.GetProductsQuantity()
	if productQuantity != 1 {
		t.Errorf("Products quantity not match!\nExpected: \"%d\"\nReceive: \"%d\"",
			1,
			productQuantity,
		)
	}
}

func TestCartTotalPrice(t *testing.T) {
	cartToTest := NewCart()
	productToBeAdd := product.NewProduct("Arroz", 3.99)

	cartToTest.AddProduct(productToBeAdd)

	totalPrice := cartToTest.GetTotalPrice()
	productToBeAddPrice := productToBeAdd.GetPrice()
	if totalPrice != productToBeAddPrice {
		t.Errorf("Total price not match!\nExpected: \"%f\"\nReceive: \"%f\"",
			productToBeAddPrice,
			totalPrice,
		)
	}
}
