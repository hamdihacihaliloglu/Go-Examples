package loops

import "fmt"

func PrimeNumberControl() {
	var value1 int
	fmt.Println("Lütfen, bir sayı giriniz:")
	fmt.Scanln(&value1)
	primeNumber := true

	if value1 > 1 {
		for i := 2; i < value1; i++ {
			if value1%i == 0 {
				primeNumber = false
			}
		}
		if primeNumber == true {
			fmt.Println("Sayı asaldır")
		} else {
			fmt.Println("Sayı asal değildir")
		}
	} else {
		fmt.Println("Girdiğiniz sayı 1 den büyük olmalıdır.")
	}

}
