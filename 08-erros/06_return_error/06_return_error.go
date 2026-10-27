package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("coffee_orders.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error opening file:", err)
		}
		return
	}
	defer file.Close()

	fmt.Println("Successfully opened file:", file.Name())

}
