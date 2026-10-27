package main

import "fmt"

func main() {

	// untyped constant with value integer
	const rewardPoints = 10

	fmt.Printf("Default type of rewardPoints: %T\n", rewardPoints) // int

	var totalRewardPoints float64 = 150.3

	// adding untyped constant to a variable of type float64, rewardPoints is implicitly converted to float64
	totalRewardPoints += rewardPoints // Implicit type conversion

	fmt.Printf("Total reward points: %.2f\n", totalRewardPoints)

}
