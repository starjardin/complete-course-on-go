package main

import "fmt"

func main() {

	// price of a cup of coffee
	price := 3.50 // flaot64

	// cups sold in one day
	quantity := 15 // int

	// total income
	// total := price * quantity // compilation error: cannot use quantity (type int) as type float64 in multiplication
	total := price * float64(quantity) // Explicit type conversion

	fmt.Printf("Total income during the day: $%.2f \n", total)

	total = float64(int(price) * quantity) // Explicit type conversion

	fmt.Printf("Total income during the day: $%.2f \n", total)

}
