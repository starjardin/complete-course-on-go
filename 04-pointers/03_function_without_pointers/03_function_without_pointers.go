package main

import "fmt"

func calculatePriceAfterDiscount(originalPrice float64, discountRate float64) float64 {
	discountAmount := originalPrice * discountRate
	return originalPrice - discountAmount
}

func main() {
	var coffeePrice float64 = 5.00
	var discount float64 = 0.10

	fmt.Println("Basic coffe price", coffeePrice)

	coffeePrice = calculatePriceAfterDiscount(coffeePrice, discount)
	fmt.Println("Coffe price after discount", coffeePrice)
}
