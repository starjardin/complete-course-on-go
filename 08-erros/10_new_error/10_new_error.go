package main

import (
	"errors"
	"fmt"
)

func CheckTemperature(temp int) error {
	if temp < 0 {
		return errors.New("temperature is too low")
	} else if temp > 100 {
		return fmt.Errorf("temperature %d is too high", temp)
	}
	return nil
}

func main() {
	temp := []int{-5, 25, 150}

	for _, t := range temp {
		err := CheckTemperature(t)
		if err != nil {
			fmt.Printf("Error checking temperature %d: %v\n", t, err)
		} else {
			fmt.Printf("Temperature %d is within the acceptable range.\n", t)
		}
	}
}
