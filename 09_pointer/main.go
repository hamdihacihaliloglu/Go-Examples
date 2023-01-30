package main

import (
	"fmt"
	"golessons/pointers"
)

func main() {
	number1 := 25
	pointers.Demo1(&number1)
	fmt.Println(number1)

	value1 := 7
	fmt.Printf("Value: %d \n", value1)
	fmt.Println("Address :", &value1)

	pointers.DemoPointers()

	numbers := []int{12, 34, 78, 89}
	fmt.Println("Main Numbers1", numbers)
	pointers.Numbers(numbers)
	fmt.Println("Main Numbers2:", numbers)
}
