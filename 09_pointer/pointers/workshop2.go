package pointers

import "fmt"

func DemoPointers() {
	var pointerNum1 *int
	num1 := 98

	fmt.Println("Address num1 :", &num1)
	fmt.Println("Value num1:", num1)

	pointerNum1 = &num1

	fmt.Println("Address pointerNum1:", &pointerNum1)
	fmt.Println("Value pointerNum1:", *pointerNum1)

	num1 = 55

	fmt.Println("Address pointerNum1:", &pointerNum1)
	fmt.Println("Value pointerNum1:", *pointerNum1)

	*pointerNum1 = 34
	fmt.Println("Address num1 :", &num1)
	fmt.Println("Value num1:", num1)

}
