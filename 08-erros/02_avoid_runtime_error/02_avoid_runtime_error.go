package main

import "fmt"

type CoffeeType string

func DispenseCoffee(coffeeAmcount int, cups int) {
	fmt.Printf("Dispensing %d grams of coffee into %d cups\n", coffeeAmcount, cups)

	if cups == 0 {
		fmt.Println("Error: Cannot dispense coffee into zero cups. Please provide a valid number of cups.")
		return
	}

	amountPerCup := coffeeAmcount / cups

	fmt.Printf("Each cup will receive %d grams of coffee\n", amountPerCup)
}

func main() {
	fmt.Println("Starting coffee machine...")
	DispenseCoffee(24, 4) // Valid case

	fmt.Println("Coffee machine is running smoothly...")
	DispenseCoffee(24, 0) // panic: runtime error: integer divide by zero
}
