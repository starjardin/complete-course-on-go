package main

type CoffeMachine interface {
	Brew() string
}

type CapsuleMachine struct {
	Brand string
}

func (c CapsuleMachine) Brew() string {
	return "Brewing a capsule coffee with " + c.Brand
}

func main() {
	var machine CoffeMachine

	machine = CapsuleMachine{Brand: "Nespresso"}
	println(machine.Brew())
}
