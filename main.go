package main

import (
	"github.com/joaovictornsv/shopping-cart-go/cart"
	"github.com/joaovictornsv/shopping-cart-go/product"
)

func main() {
	cart := cart.NewCart()

	cart.AddProduct(product.NewProduct("Arroz", 13.21))
	cart.AddProduct(product.NewProduct("Feijao", 10.98))
	cart.AddProduct(product.NewProduct("Biscoito", 4.56))
	cart.AddProduct(product.NewProduct("Sorvete", 7.54))

	cart.ShowProducts()

}
