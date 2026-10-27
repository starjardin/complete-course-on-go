package main

import (
	"fmt"
)

func main() {
	desserts := [3]string{"Cake", "Pie", "Ice Cream"}

	slice := desserts[:1]

	fmt.Println("Desserts:", desserts)
	fmt.Println("Slice:", slice)
	fmt.Println("Length", len(slice), "Capacity", cap(slice))
	slice = append(slice, "Donut")

	fmt.Println("Desserts after append:", desserts)
	fmt.Println("Slice after append:", slice)
	fmt.Println("Length", len(slice), "Capacity", cap(slice))

	slice = append(slice, "Cupcake")
	fmt.Println("Desserts after append:", desserts)
	fmt.Println("Slice after append:", slice)
	fmt.Println("Length", len(slice), "Capacity", cap(slice))

	slice = append(slice, "Brownie")
	fmt.Println("Desserts after append:", desserts)
	fmt.Println("Slice after append:", slice)
	fmt.Println("Length", len(slice), "Capacity", cap(slice))

	slice = append(slice, "Cookie")
	fmt.Println("Desserts after append:", desserts)
	fmt.Println("Slice after append:", slice)
	fmt.Println("Length", len(slice), "Capacity", cap(slice))

}
