package main

import "fmt"

type CoffeeType string

func DispenseCoffee(coffeeAmcount int, cups int) {
	fmt.Printf("Dispensing %d grams of coffee into %d cups\n", coffeeAmcount, cups)

	amountPerCup := coffeeAmcount / cups

	fmt.Printf("Each cup will receive %d grams of coffee\n", amountPerCup)
}

func main() {
	fmt.Println("Starting coffee machine...")
	DispenseCoffee(24, 0) // panic: runtime error: integer divide by zero

	fmt.Println("Coffee machine is running smoothly...")
}
