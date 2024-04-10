package main

import "fmt"

func main() {
	transactions := [5]int{1, 2, 3, 4, 5}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[0] = "Тинькофф"
	fmt.Println(banks)
	patrial := transactions[1:]
	fmt.Println(patrial)
}
