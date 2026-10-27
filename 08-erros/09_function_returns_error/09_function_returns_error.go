package main

import (
	"fmt"
)

type OutOfStockError struct {
	Item string
}

func (e OutOfStockError) Error() string {
	return fmt.Sprintf("%s is out of stock", e.Item)
}

func ServedDrink(item string) (string, error) {
	stockCount, exists := stock[item]
	if !exists {
		return "", fmt.Errorf("item %s does not exist", item)
	}
	if stockCount <= 0 {
		return "", OutOfStockError{Item: item}
	}
	stock[item]--
	return fmt.Sprintf("Here is your %s!", item), nil
}

var stock = map[string]int{
	"Espresso":   10,
	"Latte":      5,
	"Cappuccino": 0,
}

func main() {
	drink, err := ServedDrink("Cappuccino")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(drink)
	}

	drink, err = ServedDrink("Mocha")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(drink)
	}

	drink, err = ServedDrink("Espresso")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(drink)
	}
}
