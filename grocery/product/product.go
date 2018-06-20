package product

import "errors"

// Product defines a product
type Product struct {
	Name    string
	Price   float32
	InStock int
}

// New creates a new product.
func New(name string, price float32, inStock int) Product {
	p := Product{name, price, inStock}
	return p
}

// IsAvailableInStock verifies if the product is available in stock.
func (product *Product) IsAvailableInStock() (available bool, err error) {
	available = product.InStock > 0
	if available {
		err = nil
	} else {
		err = errors.New(product.Name + " is out of stock")
	}

	return
}
