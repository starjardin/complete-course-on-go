package main

import "fmt"

func main() {
	var coffeName = "Espresso"
	var size = "Small"
	var price = 2.50

	fmt.Println(size, coffeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f \n", size, coffeName, price)
}
