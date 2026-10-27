package main

import (
	"fmt"
)

func main() {
	ratings := []int{5, 4, 3, 4, 5} // declare a slice of 5 integers

	fmt.Println("Original Ratings:", ratings)

	ratings[2] = 5 // modify the third element of the slice

	fmt.Println("Updated element with index 2:", ratings)
	fmt.Println("Length of the slice is:", len(ratings))

	coffeeTypes := make([]string, 3) // create a slice of strings with length 3

	fmt.Println("Coffee Types slice after make():", coffeeTypes)

	coffeeTypes[0] = "Espresso"
	coffeeTypes[1] = "Latte"
	coffeeTypes[2] = "Cappuccino"

	fmt.Println("Updated Coffee Types slice after reassignment:", coffeeTypes)
}
