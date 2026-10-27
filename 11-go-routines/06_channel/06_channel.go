package main

import (
	"fmt"
	"time"
)

func prepareDrink(orderChannel chan string) {
	fmt.Printf("Barista : Starting to make coffee\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista : Done\n")
	orderChannel <- "Latte"
}

func main() {
	orderChannel := make(chan string)
	fmt.Println("Coffe shop opens")

	go prepareDrink(orderChannel)

	coffee := <-orderChannel
	fmt.Printf("Customer received: %s\n", coffee)

	fmt.Println("All drinks are ready")
	fmt.Println("Coffe shop closes")

}
