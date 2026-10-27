package main

import (
	"fmt"
)

func main() {
	stock := make(map[string]int) // initialize the map

	stock["Espresso"] = 10
	stock["Capuccino"] = 5
	fmt.Println("Product in stock:", stock)
}
