package main

import "fmt"

func main() {

	// var shopName = "Coffee House"

	// shopName = "Coffee House 2.0"

	// fmt.Println("Welcome to", shopName)

	const shopName = "Coffee House"

	// const shopName string // compile-time error: shopName has to be initialized

	// shopName = "Coffee House 2.0" // This will cause a compile-time error because shopName is a constant

	fmt.Println("Welcome to", shopName)

}
