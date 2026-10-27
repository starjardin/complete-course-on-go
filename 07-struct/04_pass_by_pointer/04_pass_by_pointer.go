package main

import "fmt"

type CoffeMachine struct {
	Model          string // default value is empty string ""
	Status         string // default value is empty string ""
	OperationHours int    // default value is 0
}

func markAsOutOfService(machine *CoffeMachine) {
	machine.Status = "Out of Service"
	fmt.Printf("Machine %s is now %s\n", machine.Model, machine.Status)
}

func main() {
	myMachine := CoffeMachine{"Model X", "Operational", 1000}
	markAsOutOfService(&myMachine)

	fmt.Printf("After markAsOutOfService, Machine %s status: %s\n", myMachine.Model, myMachine.Status)
}
