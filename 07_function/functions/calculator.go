package functions

func SumNumbers(numbers ...int) int {
	sumNumbers := 0
	for i := 0; i < len(numbers); i++ {
		sumNumbers = sumNumbers + numbers[i]
	}
	return sumNumbers
}

func Calculator(num1 int, num2 int) (int, int, int, float32) {
	add := num1 + num2
	subtract := num1 - num2
	multiply := num1 * num2
	divide := float32(num1 / num2)

	return add, subtract, multiply, divide

}
