package main

import (
	"fmt"
	"./grocery"
)

func main() {
	var egg = grocery.NewProduct("Egg", 2.0, 50)
	var milk = grocery.NewProduct("Milk", 5.0, 20)
	var outOfStock = grocery.NewProduct("Out of Stock", 100000000, 0)
	
	var customer = grocery.NewCustomer()

	customer.AddToCart(egg)
    customer.AddToCart(egg)
    customer.AddToCart(egg)
    customer.AddToCart(milk)
    customer.AddToCart(outOfStock)

	fmt.Printf("Total: $%v\n", customer.Buy())
	fmt.Printf("Total (Empty Cart): $%v\n", customer.Buy())
}
