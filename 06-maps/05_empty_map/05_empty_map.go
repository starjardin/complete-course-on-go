package main

import (
	"fmt"
)

func main() {
	var stock map[string]int // zero value is nil

	stock["Espresso"] = 10

	fmt.Println("Stock:", stock)

	if stock == nil {
		fmt.Println("Stock map is nil")
	}

}
