package main

import "fmt"

func createTemperatureAdjuster() (func(float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemp := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}
	return adjustTemp, baseTemperature
}

func main() {

	adjustTemp, baseTemp := createTemperatureAdjuster()

	fmt.Printf("Base temperature: %.2f degrees\n", baseTemp)

	fmt.Printf("Adjusted temp +1.5: %.2f degrees\n", adjustTemp(1.5))
	fmt.Printf("Adjusted temp -3.0: %.2f degrees\n", adjustTemp(-3.0))
	fmt.Printf("Adjusted temp +5.0: %.2f degrees\n", adjustTemp(5.0))
}
