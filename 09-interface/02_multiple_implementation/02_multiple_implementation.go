package main

import "fmt"

type CoffeMachine interface {
	Brew() string
}

type CapsuleMachine struct {
	Brand string
	Model string
	Price int
}

type DripeMachine struct {
	Model string
	Price int
}

func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s %s has brewed a cup of capsule coffee", c.Brand, c.Model)
}

func (c DripeMachine) Brew() string {
	return fmt.Sprintf("%s has brewed a cup of drip coffee", c.Model)
}

func (c CapsuleMachine) DeepClean() {
	fmt.Printf("%s %s is being deep cleaned\n", c.Brand, c.Model)
}

func main() {
	var machineOne, machineTwo CoffeMachine

	machineOne = CapsuleMachine{Brand: "Nespresso", Model: "Vertuo", Price: 150}
	println(machineOne.Brew())

	machineTwo = DripeMachine{Model: "Hario V60", Price: 100}
	println(machineTwo.Brew())
}
