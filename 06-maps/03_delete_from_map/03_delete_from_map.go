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

	fmt.Println("Initial menu:", menu)

	delete(menu, "Latte")
	fmt.Println("Menu after deleting Latte:", menu)

	delete(menu, "Tea") // nothing will happen if the key doesn't exist
	fmt.Println("Menu after trying to delete Tea:", menu)

}
