package main

import "fmt"

func main() {
	var cupsQty int = 3

	// cupsQty = true # compilation error

	fmt.Printf("Cups quantity: %d\n", cupsQty)

	var wasProcessed = true

	// wasProcessed = 42 # compilation error

	fmt.Printf("Was processed: %t\n", wasProcessed)

}
