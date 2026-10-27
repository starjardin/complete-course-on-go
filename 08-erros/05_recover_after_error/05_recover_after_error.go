package main

import "fmt"

type CoffeeType string

func DispenseCoffee(coffeeAmcount int, cups int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Machine error:", r)
		}
	}()

	fmt.Printf("Dispensing %d grams of coffee into %d cups\n", coffeeAmcount, cups)

	amountPerCup := coffeeAmcount / cups

	fmt.Printf("Each cup will receive %d grams of coffee\n", amountPerCup)
}

func main() {
	fmt.Println("Starting coffee machine...")
	DispenseCoffee(24, 4) // Valid case

	fmt.Println("C\noffee machine is running smoothly...")
	DispenseCoffee(24, 0) // we recover from this panic and the program continues to run

	fmt.Println("\nCoffee machine is still running smoothly...")

	DispenseCoffee(500, 150)
}
