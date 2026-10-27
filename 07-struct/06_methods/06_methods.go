package main

import "fmt"

type CoffeeShop struct {
	Model string // default value is empty string ""
}

// method with value receiver
func (shop CoffeeShop) Greet() {
	fmt.Printf("Welcome to %s!\n", shop.Model)
}

func main() {
	myShop := CoffeeShop{
		Model: "Brew Haven",
	}

	myShop.Greet()
}
