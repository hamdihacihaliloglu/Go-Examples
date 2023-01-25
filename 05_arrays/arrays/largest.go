package arrays

import "fmt"

func LargestNumber() {
	numbers := [5]int{35, 45, 3, 67, 45}

	largestnumber := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > largestnumber {
			largestnumber = numbers[i]
		}
	}
	fmt.Println(largestnumber)
}
