package structs

import "fmt"

type product struct {
	name      string
	unitPrice float64
	brand     string
	category  string
}

func Workshop1() {
	fmt.Println(product{"Pulse", 350.00, "MSI", "Computer"})
	fmt.Println(product{name: "Pulse", unitPrice: 450.50})
}
