package main

import (
	"fmt"
)

func main() {
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"} // declare an array of 4 strings

	slice := dessertMenu[1:2] // create a slice from the array (includes index 1 and 2)

	fmt.Println("Dessert Menu:", dessertMenu)
	fmt.Println("Slice of Dessert Menu:", slice)

	slice = dessertMenu[:] // create a slice that includes all elements of the array

	fmt.Println("Slice of entire Dessert Menu:", slice)
}
