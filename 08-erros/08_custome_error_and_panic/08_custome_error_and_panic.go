package main

import (
	"fmt"
)

type CoffeeError string

func (e CoffeeError) Error() string {
	return string(e)
}

func main() {

	var err error
	err = CoffeeError("Out of coffee beans")

	if err != nil {
		fmt.Println("An error occurred:", err)
		panic(err)
	} else {
		fmt.Println("There is no error")
	}

}
