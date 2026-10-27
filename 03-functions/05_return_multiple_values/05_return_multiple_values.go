package main

import "fmt"

func processessPayment(orderTotal float64, tip float64, amountPayd float64) (float64, float64) {
	totalAmountDue := orderTotal + tip
	change := amountPayd - totalAmountDue

	return totalAmountDue, change
}

func main() {
	totalCost, change := processessPayment(6.50, 2.00, 10.00)

	fmt.Printf("Total cost (with tip): $%.2f\n", totalCost)
	fmt.Printf("Change returned to the customer: $%.2f\n", change)
}
