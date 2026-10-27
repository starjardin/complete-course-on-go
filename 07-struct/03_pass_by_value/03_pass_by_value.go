package main

import "fmt"

/* List of types in Go which are passed by value:**
1. Basic types: int, float64, bool, string, etc.
2. Structs: Custom data types defined using the struct keyword.
3. Arrays: Fixed-size collections of elements of the same type.
4. Pointers: Memory addresses of variables.
*/

type CoffeOrder struct {
	CoffeType string // default value is empty string ""
	Size      string // default value is empty string ""
}

func printReceiptCopy(order CoffeOrder) {
	order.CoffeType = "Changed Type"
	order.Size = "Changed Size"
	fmt.Printf("Changed order: %s - %s\n", order.CoffeType, order.Size)
}

func main() {
	myOrder := CoffeOrder{"Cappuccino", "Medium"}
	fmt.Printf("Original Order: %s - %s\n", myOrder.CoffeType, myOrder.Size)

	printReceiptCopy(myOrder)

	fmt.Printf("After printReceiptCopy, Original Order: %s - %s\n", myOrder.CoffeType, myOrder.Size)
}
