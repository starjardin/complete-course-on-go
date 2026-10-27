package main

import "fmt"

type CoffeeOrder struct {
	CustomerName string
	CoffeeSize   string
	CoffeeType   string
}

func main() {

	//Struct version
	orderStruct := CoffeeOrder{
		CustomerName: "John Doe",
		CoffeeSize:   "Large",
		CoffeeType:   "Espresso",
	}

	// Map version
	orderMap := map[string]string{
		"CustomerName": "John Doe",
		"CoffeeSize":   "Large",
		"CoffeeType":   "Espresso",
	}

	fmt.Println("--------Using strtuct--------")
	fmt.Printf("Customer Name: %s\n", orderStruct.CustomerName)
	fmt.Printf("Coffee Size: %s\n", orderStruct.CoffeeSize)
	fmt.Printf("Coffee Type: %s\n", orderStruct.CoffeeType)

	fmt.Println("--------Using map--------")
	fmt.Printf("Customer Name: %s\n", orderMap["CustomerName"])
	fmt.Printf("Coffee Size: %s\n", orderMap["CoffeeSize"])
	fmt.Printf("Coffee Type: %s\n", orderMap["CoffeeType"])
}
