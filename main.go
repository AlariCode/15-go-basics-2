package main

import "fmt"

func main() {
	transactions := []int{0, 20, 35}
	temp := transactions
	transactions = append(transactions, 100)
	fmt.Println(temp)
	fmt.Println(transactions[1:])
}
