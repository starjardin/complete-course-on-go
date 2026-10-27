package main

import (
	"fmt"
)

func sellCoffee(stock map[string]int, coffeeType string, quantity int) {
	currentStock, exist := stock[coffeeType]
	if !exist {
		fmt.Printf("Sorry, we don't have %s in stock.\n", coffeeType)
		return
	}

	if currentStock >= quantity {
		stock[coffeeType] = currentStock - quantity
		fmt.Printf("Sold %d %s. Remaining stock: %d\n", quantity, coffeeType, stock[coffeeType])
	} else {
		fmt.Printf("Sorry, we only have %d %s in stock. Cannot sell %d.\n", currentStock, coffeeType, quantity)
	}
}

func main() {
	coffeeStock := map[string]int{
		"Espresso":  5,
		"Capuccino": 2,
		"Latte":     1,
		"Americano": 3,
	}

	sellCoffee(coffeeStock, "Espresso", 1)
	sellCoffee(coffeeStock, "Capuccino", 10)
	sellCoffee(coffeeStock, "Tea", 10)

}
