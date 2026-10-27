package main

import "fmt"

type Barista interface {
	PrepareCoffee() string
}

type SeniorBarista struct {
	Name string
}

type JuniorBarista struct {
	Name string
}

func (s SeniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s has prepared a cup of coffee with expert skills", s.Name)
}

func (j JuniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s has prepared a cup of coffee with basic skills", j.Name)
}

func ServeDrink(b Barista) {
	fmt.Println(b.PrepareCoffee())
	fmt.Printf("Barista served coffe to the client\n\n")
}

func main() {
	tantely := SeniorBarista{Name: "Tantely"}
	var joel Barista = JuniorBarista{Name: "Joel"}

	ServeDrink(tantely)
	ServeDrink(joel)

	joel = SeniorBarista{Name: "Joel"}
	ServeDrink(joel)
}
