package main

import "fmt"

func greet(coffeeShopName string) {
	fmt.Printf("Welcome to the coffee shop %s\n", coffeeShopName)
}

func main() {
	greet("Brew & Beans")
	greet("Coffee & Milk")
}
