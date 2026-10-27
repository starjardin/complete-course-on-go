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

	for drink, price := range menu {
		fmt.Printf("%s: $%.2f\n", drink, price)
	}

}
