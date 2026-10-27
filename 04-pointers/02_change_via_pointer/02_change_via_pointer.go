package main

import "fmt"

func main() {
	var coffeePrice = 4.50

	fmt.Println("Coffe price", coffeePrice)

	// STEP 1

	// compile Time (code you write) : var coffeePrice = 4.50
	// run Time (what machine sees) : [some memory address] = 4.50

	// STEP 2
	// compile Time (code you write) : fmt.Println("Coffe price", coffeePrice)
	// run Time (what machine sees) : fmt.Println([some memory address], [some memory address] = 4.50)

	fmt.Println("Memory address of coffeePrice 4.50", &coffeePrice)

	coffeePrice = 5.00

	fmt.Println("Memory address of coffeePrice 5.00", &coffeePrice)

	var pointerToCoffeePrice *float64 = &coffeePrice

	/*
		go to the memory where pointerToCoffeePrice
		points to and change the value there to 5.50
	*/
	*pointerToCoffeePrice = 7.50

	fmt.Println("Coffe price after change via pointer", coffeePrice)
}
