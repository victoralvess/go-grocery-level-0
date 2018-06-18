package grocery

import (
	"fmt"
)

type Customer = customer
type Product = product

type customer struct {
	cart []Product
}

func NewCustomer() Customer {
	return Customer {}
}

func (c *Customer) AddToCart(product Product) {
	if product.InStock <= 0 {
		fmt.Printf("%s is out of stock\n", product.Name)
		return
	}
	
	product.InStock -= 1
	c.cart = append(c.cart, product)
}

func (c *Customer) Buy() float32 {
	var total float32
	for i := 0; i < len(c.cart); i++ {
		total += c.cart[i].Price
	}

	c.EmptyCart()

	return total
}

func (c *Customer) EmptyCart() {
	c.cart = nil
}
