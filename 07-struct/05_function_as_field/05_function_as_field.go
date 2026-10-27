package main

import "fmt"

type CoffeeShop struct {
	Model string                // default value is empty string ""
	Greet func(shop CoffeeShop) // function field to greet customers
}

func greetShop(shop CoffeeShop) {
	fmt.Printf("Welcome to %s!\n", shop.Model)
}

func main() {
	myShop := CoffeeShop{
		Model: "Brew Haven",
		Greet: greetShop,
	}

	myShop.Greet(myShop)
}
