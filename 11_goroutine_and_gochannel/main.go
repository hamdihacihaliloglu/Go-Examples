package main

import (
	"fmt"
	"golessons/routinechannel"
)

func main() {
	IncomeCn := make(chan int)
	ExpensesCn := make(chan int)

	go routinechannel.Income(IncomeCn)
	go routinechannel.Expenses(ExpensesCn)

	incomeResult, expensesResult := <-IncomeCn, <-ExpensesCn
	result := incomeResult - expensesResult
	fmt.Println("Account : ", result)
}
