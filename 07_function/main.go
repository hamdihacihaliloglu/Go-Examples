package main

import (
	"fmt"
	"golessons/functions"
)

func main() {
	result := functions.SumNumbers(12, 12, 32, 54)
	fmt.Println(result)

	result2 := []int{25, 32, 45, 456}
	fmt.Println(functions.SumNumbers(result2...))

	resultAdd, resultSubtract, resultMultiply, resultDivide := functions.Calculator(35, 4)
	fmt.Println("Add:", resultAdd)
	fmt.Println("Subtract:", resultSubtract)
	fmt.Println("Multiply:", resultMultiply)
	fmt.Println("Divide:", resultDivide)

}
