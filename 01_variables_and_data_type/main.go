package main

import "fmt"

func main() {
	var name string = "Hamdi"
	fmt.Println(name)

	var number int = 15
	fmt.Println(number)

	var number2 float32 = 0.18
	fmt.Println(number2)
	fmt.Println(100 + 100*number2)

	number3 := 25.01
	fmt.Println(number3)
	fmt.Printf("Data Type: %T \n", number3)

	var status bool

	var text1 string = "Hamdi"
	var text2 string = "Anil"

	status = text1 == text2
	fmt.Println(status) //output:false

	status = text1 != text2
	fmt.Println(status) //output:true

	fmt.Println(!status) //output:false

}
