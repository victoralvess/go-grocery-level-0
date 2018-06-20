package customer

import (
	"fmt"

	"github.com/victoralvess/go-grocery-level-0/grocery/product"
)

// Customer defines a customer
type Customer struct {
	cart []product.Product
}

// New creates a new customer.
func New() Customer {
	return Customer{}
}

// AddToCart adds a product to the customer cart if the product is available in stock.
func (c *Customer) AddToCart(product product.Product) {
	if available, err := product.IsAvailableInStock(); available && err == nil {
		product.InStock--
		c.cart = append(c.cart, product)
	} else {
		fmt.Println(err)
	}
}

// Buy buys all items from customer's cart.
func (c *Customer) Buy() float32 {
	var total float32
	for i := 0; i < len(c.cart); i++ {
		total += c.cart[i].Price
	}

	c.EmptyCart()

	return total
}

// EmptyCart removes all items from customer's cart.
func (c *Customer) EmptyCart() {
	c.cart = nil
}
