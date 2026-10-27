package main

import "fmt"

func applyDiscount(price *float64, discountRate float64) {
	*price = *price - (*price * discountRate)
}

func main() {
	var coffeePrice float64 = 5.00
	var discount float64 = 0.10

	fmt.Printf("Basic coffe price %.2f\n", coffeePrice)

	applyDiscount(&coffeePrice, discount)
	fmt.Printf("Coffe price after discount %.2f\n", coffeePrice)
}
