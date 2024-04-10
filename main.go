package main

import "fmt"

func main() {
	transactions := [3]int{1, 2, 3}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[0] = "Тинькофф"
	fmt.Println(banks)
}
