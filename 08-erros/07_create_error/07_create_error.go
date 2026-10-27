package main

import (
	"fmt"
)

func main() {
	var err error

	err = fmt.Errorf("an example error occurred")

	// err = "Some other error" // This will cause a compile-time error because err is of type error, not string

	if err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		fmt.Println("There is no error")
	}

}
