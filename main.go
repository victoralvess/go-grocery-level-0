package main

import (
	"fmt"
	"github.com/victoralvess/go-grocery-level-0/grocery/product"
	"github.com/victoralvess/go-grocery-level-0/grocery/customer"
)

func main() {
	var egg = product.New("Egg", 2.0, 50)
	var milk = product.New("Milk", 5.0, 20)
	var outOfStock = product.New("Out of Stock", 100000000, 0)
	
	var customer = customer.New()

	customer.AddToCart(egg)
    customer.AddToCart(egg)
    customer.AddToCart(egg)
    customer.AddToCart(milk)
    customer.AddToCart(outOfStock)

	fmt.Printf("Total: $%v\n", customer.Buy())
	fmt.Printf("Total (Empty Cart): $%v\n", customer.Buy())
}
