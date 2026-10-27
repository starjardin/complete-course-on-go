package main

import "fmt"

func estimatedBrewTime(cupsQty int, secondsPercup int) (totalTimeSeconds int) {
	totalTimeSeconds = cupsQty * secondsPercup
	return // naked return
}

// func estimatedBrewTime(cupsQty int, secondsPercup int) int {
// 	totalTimeSeconds := cupsQty * secondsPercup

// 	return totalTimeSeconds
// }

func main() {
	brewTime := estimatedBrewTime(12, 20)
	fmt.Printf("Estimated brew time: %d seconds\n", brewTime)
}
