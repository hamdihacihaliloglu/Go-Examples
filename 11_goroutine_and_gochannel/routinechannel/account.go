package routinechannel

func Income(IncomeCn chan int) {
	income := []int{1200, 3500, 350, 225}
	sum := 0
	for i := 0; i < len(income); i++ {
		sum = sum + income[i]
	}
	IncomeCn <- sum
}

func Expenses(ExpensesCn chan int) {
	expenses := []int{300, 250, 126}
	sum := 0
	for i := 0; i < len(expenses); i++ {
		sum = sum + expenses[i]
	}
	ExpensesCn <- sum
}
