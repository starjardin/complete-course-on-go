package main

import (
	"fmt"
)

func main() {
	menu := [3]string{"Tea", "Coffee", "Juice"} // declare an array of 3 strings

	slice := menu[:2] // create a slice from the array (includes index 0 and 1)

	fmt.Println("Before slice change")
	fmt.Println("Menu:", menu)
	fmt.Println("Slice of Menu:", slice)

	fmt.Println("Length of menu array:", len(menu))

	slice[0] = "Matcha" // modify the first element of the slice

	fmt.Println("After slice change")
	fmt.Println("Menu:", menu) // also changed
	fmt.Println("Slice of Menu:", slice)
}
