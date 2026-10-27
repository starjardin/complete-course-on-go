package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"Espresso": 5.99,
	}

	drink := "Capuccino"

	fmt.Printf("%s price is $%.2f\n", drink, menu[drink])

	price, exist := menu[drink]
	if exist {
		fmt.Printf("%s price is $%.2f\n", drink, price)
	} else {
		fmt.Printf("%s is not on the menu\n", drink)
	}

}
