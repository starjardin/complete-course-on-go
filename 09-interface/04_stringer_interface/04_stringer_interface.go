package main

import "fmt"

type Order struct {
	Customer string
	Item     string
	Quantity int
}

func (o Order) String() string {
	return fmt.Sprintf("Customer: %s, Item: %s, Quantity: %d", o.Customer, o.Item, o.Quantity)
}

func main() {
	order := Order{Customer: "Alice", Item: "Coffee", Quantity: 2}
	fmt.Println(order)
}
