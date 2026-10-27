package main

import "fmt"

type CoffeOrder struct {
	CoffeType string // default value is empty string ""
	Size      string // default value is empty string ""
}

func main() {

	// struct literal
	order := CoffeOrder{
		CoffeType: "Espresso",
		Size:      "Medium",
	}

	fmt.Println("Order", order)

	order.CoffeType = "Cappuccino"
	order.Size = "Small"

	fmt.Println("Order", order)

	// inline struct literal
	myFriendOrder := CoffeOrder{"Latte", "Large"}
	fmt.Println("My friend's order type", myFriendOrder.CoffeType)
	fmt.Println("My friend's order size", myFriendOrder.Size)
}
