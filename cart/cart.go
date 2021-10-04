package cart

import (
	"fmt"

	"github.com/joaovictornsv/shopping-cart-go/product"
)

type Cart struct {
	productsList []product.Product
}

func NewCart() Cart {
	return Cart{productsList: make([]product.Product, 0, 20)}
}

func (c *Cart) AddProduct(product product.Product) {
	c.productsList = append(c.productsList, product)
}

func (c *Cart) GetTotalPrice() float64 {
	var total float64 = 0

	for i := 0; i < len(c.productsList); i++ {
		total += c.productsList[i].GetPrice()
	}

	return total
}

func (c *Cart) ShowProducts() {
	fmt.Println("========== Products ==========")
	for i := 0; i < len(c.productsList); i++ {
		fmt.Printf("%-10s R$%.2f\n", c.productsList[i].GetName(), c.productsList[i].GetPrice())
	}
	fmt.Println("------------------------------")
	fmt.Printf("Your cart have %v products\n", c.GetProductsQuantity())
	fmt.Printf("%-10s R$%.2f\n", "Total:", c.GetTotalPrice())
	fmt.Println("==============================")
}

func (c *Cart) GetProductsQuantity() int {
	return len(c.productsList)
}
