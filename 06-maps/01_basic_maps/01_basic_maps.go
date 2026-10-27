package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"Asspresso": 5.99,
		"Capuccino": 2.99,
		"Latte":     1.49,
		"Americano": 3.49,
	}

	fmt.Println("Menu:", menu)

	// Accessing a value by key
	price := menu["Latte"]
	fmt.Println("Price of Latte:", price)

	menu["Latte"] = 5.49
	fmt.Println("New Latte price:", menu["Latte"])

	menu["Mocha"] = 4.99
	fmt.Println("Menu after adding Mocha:", menu)

}
