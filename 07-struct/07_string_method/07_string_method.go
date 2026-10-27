package main

import "fmt"

type CoffeeType string

func (s CoffeeType) Describe() {
	fmt.Printf("The string is: %s\n", s)
}

func main() {
	myCoffee := CoffeeType("Espresso")
	myCoffee.Describe()
}
