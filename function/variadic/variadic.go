package main

import (
	"fmt"
)

func sum(num ...int) int {
	var result int = 0
	for _, number := range num {
		result += number
	}
	return result
}

func Buy(name string, spend ...float64) (string, float64) {
	var money float64 = 0
	for _, x := range spend {
		money += x
	}
	return name, money
}

func main() {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(sum(num...))
	name, spend := Buy("John", float64(num[5]))
	// fmt.Sprintf() cannot be printing directly
	result := fmt.Sprintf("Name: %s, spend %f on items", name, spend)
	fmt.Print(result)
}
