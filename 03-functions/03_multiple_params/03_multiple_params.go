package main

import "fmt"

func getDrinkInfo(customerName string, drinkType string) {
	info := "%s's favorite drink is %s\n"
	fmt.Printf(info, customerName, drinkType)
}

func main() {
	getDrinkInfo("Alice", "Latte")
	getDrinkInfo("Bob", "Cappuccino")
}
