package main

import (
	"fmt"
)

func barista() {
	fmt.Println("barista: Starting to make coffee")
	fmt.Println("barista: Done")
}

func main() {

	fmt.Println("Coffe shop opens")

	go barista()

	fmt.Println("Coffe shop closes")

}
