package loops

import (
	"fmt"
	"math"
)

func BigOn(n int) {
	var i int
	for i = 0; i <= n; i++ {
		fmt.Println(i)
	}
}

func BigOn2(n int) {
	var i, j int
	for i = 0; i <= n; i++ {
		for j = 0; j <= n; j++ {
			fmt.Print(i)
			fmt.Println(j)
		}
	}
}

func BigOn3(n int) {
	var i, j, k int
	for i = 0; i <= n; i++ {
		for j = 0; j <= n; j++ {
			for k = 0; k <= n; k++ {
				fmt.Print(i)
				fmt.Print(j)
				fmt.Println(k)
			}
		}
	}
}

func Logn(n float64) {
	for n > 1 {
		n = math.Floor(n / 2)
		fmt.Println(n)
	}
}

func NLogn(n float64) {
	var lim int
	lim = int(n)
	i := 1
	for n > 1 {
		n = math.Floor(n / 2)
		for i = 1; i <= lim; i++ {
			fmt.Println(i)
		}
	}
}
