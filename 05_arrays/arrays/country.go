package arrays

import "fmt"

func Citys() {
	var citys [5]string
	citys[0] = "İstanbul"
	citys[1] = "Ankara"
	citys[2] = "Samsun"
	citys[3] = "Adana"
	citys[4] = "Kars"

	for i := 0; i < 5; i++ {
		fmt.Println(citys[i])
	}
}
