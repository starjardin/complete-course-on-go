package main

import (
	"fmt"
	"time"
)

func prepareDrink(barista string, drink string, c chan string) {
	fmt.Printf("Barista %s : Starting to make coffee\n", barista)
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista %s: %s is ready!\n", barista, drink)
	c <- drink
}

func main() {
	fmt.Println("Coffe shop opens...")
	c := make(chan string)

	baristas := []string{"Alice", "Bob", "Charlie"}
	drinks := []string{"Latte", "Cappuccino", "Espresso"}

	for i, name := range baristas {
		go prepareDrink(name, drinks[i], c)
	}

	for range baristas {
		<-c
	}

	fmt.Println("All drinks are ready")
	fmt.Println("Coffe shop closes")

}
