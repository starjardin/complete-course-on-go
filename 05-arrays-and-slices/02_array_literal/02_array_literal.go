package main

import "fmt"

func main() {
	coffeetypes := [3]string{"Espresso", "Latte", "Cappuccino"} // declare an array of 3 strings
	fmt.Println("Types of coffee:", coffeetypes)
	fmt.Println("Length of array", len(coffeetypes))
}
