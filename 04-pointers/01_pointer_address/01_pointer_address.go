package main

import "fmt"

func main() {
	coffee := "Espresso"

	// Get the memory address of the variable 'coffee'
	address := &coffee

	fmt.Printf("The memory address of 'coffee' is: %p\n", address)

	copyCoffee := coffee // copy value to another location

	fmt.Printf("The memory address of 'copyCoffee' is: %p\n", &copyCoffee) // defers from coffee address
}
