package conditional

import "fmt"

func WorkshopDemo() {
	var num1, num2, num3 int = 18, 12, 15
	var tempnum1 int = num1

	if num2 > tempnum1 {
		tempnum1 = num2

	}
	if num3 > tempnum1 {
		tempnum1 = num3

	}
	fmt.Println("En Büyük Sayı", tempnum1)
}
