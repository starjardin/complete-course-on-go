package main

import "fmt"

func main() {

	var coffeeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 4.25

	fmt.Printf("Order: %d %s(s) at $%.2f each\n", quantity, coffeeType, unitPrice)

	var (
		customerName string = "Alice"
		isReadyToPay bool   = true
		tableNumber  int    = 5
	)

	fmt.Printf("Customer %s at table %d is ready to pay %t\n", customerName, tableNumber, isReadyToPay)

	const (
		SizeSmall  = "S"
		SizeMedium = "M"
		SizeLarge  = "L"
	)
}
