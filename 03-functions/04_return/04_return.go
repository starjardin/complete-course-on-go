package main

import "fmt"

func calculateLayaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func updateTotalPoints(currentPoints int, pointsEarned int) int {
	return currentPoints + pointsEarned
}

func main() {
	totalPoints := 120
	newlyEarnedPoints := calculateLayaltyPoints(9.30)
	fmt.Printf("Earned points today: %d\n", newlyEarnedPoints)
	totalPoints = updateTotalPoints(totalPoints, newlyEarnedPoints)
	fmt.Printf("Total points: %d\n", totalPoints)
}
