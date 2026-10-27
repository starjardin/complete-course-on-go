package main

import "fmt"

type CoffeOrder struct {
	CoffeType    string // default value is empty string ""
	Size         string // default value is empty string ""
	CustomerName string // default value is empty string ""
	Points       int    // default value is 0
}

func main() {
	var order CoffeOrder

	fmt.Println("Order", order)

	order.CoffeType = "Cappuccino"
	order.Size = "Small"
	order.CustomerName = "John Doe"
	order.Points = 10

	fmt.Println("Order", order)
}
