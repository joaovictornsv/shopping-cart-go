package product

type Product struct {
	name  string
	price float64
}

func NewProduct(name string, price float64) Product {
	return Product{name: name, price: price}
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetPrice() float64 {
	return p.price
}
