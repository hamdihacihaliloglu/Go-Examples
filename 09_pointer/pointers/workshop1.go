package pointers

import "fmt"

func Demo1(number1 *int) int {
	*number1 = *number1 + 2
	return *number1
}

func Numbers(numbers []int) {
	numbers[2] = 26
	fmt.Println("Numbers Function:", numbers[2])

}
