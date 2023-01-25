package loops

import "fmt"

func AmicableNumbersControl() {
	var num1, num2 int
	sum1 := 0
	sum2 := 0
	fmt.Print("Birinci sayıyı giriniz:")
	fmt.Scanln(&num1)
	fmt.Print("İkinci sayısı girniz:")
	fmt.Scanln(&num2)

	for i := 1; i < num1; i++ {
		if num1%i == 0 {
			sum1 = sum1 + i
		}
	}
	for i := 1; i < num2; i++ {
		if num2%i == 0 {
			sum2 = sum2 + i
		}
	}
	if sum1 == num2 && sum2 == num1 {
		fmt.Printf("%v ve %v arkadaş sayılardır.", num1, num2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayılar değildir.", num1, num2)
	}
}
