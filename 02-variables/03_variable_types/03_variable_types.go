package main

import "fmt"

func main() {
	var stringType = "Hello, World!"
	var integerType = 42
	var floatType = 3.14
	var booleanType = true

	fmt.Printf("%T\n", stringType)
	fmt.Printf("%T\n", integerType)
	fmt.Printf("%T\n", floatType)
	fmt.Printf("%T\n", booleanType)
}
