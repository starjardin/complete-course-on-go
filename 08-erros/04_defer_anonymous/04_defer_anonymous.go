package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Closing the coffee shop...")
	}()

	defer fmt.Println("Brewing a fresh cup of espresso...")
	fmt.Println("Brewing a fresh cup of cappuccino...")
}
