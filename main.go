package main

import "fmt"

func main() {
	transactions := [6]int{1, 2, 3, 4, 5, 6}
	transactionsPartial := transactions[1:5]
	transactionsNewPartial := transactionsPartial[:1]
	transactionsNewPartial[0] = 30

	transactionsNewPartial = transactionsNewPartial[0:4]

	fmt.Println(transactions)
	fmt.Println(transactionsNewPartial)
	fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))
}
