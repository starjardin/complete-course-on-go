package main

import (
	"fmt"
)

func main() {
	menu := []string{"Cake", "pie"}

	fmt.Println("Initial menu:", menu)
	fmt.Println("Length", len(menu), "Capacity", cap(menu))

	menu2 := append(menu, "Donut")

	fmt.Println("Menu after append:", menu)
	fmt.Println("Length", len(menu2), "Capacity", cap(menu2))

	slice := make([]string, 0, 4)

	// slice[0] = "Cake" // This will cause a runtime error: index out of range

	slice = append(slice, "Cake", "Pie", "Donut")

	fmt.Println("Slice after append:", slice)
	fmt.Println("Length", len(slice), "Capacity", cap(slice))
}
