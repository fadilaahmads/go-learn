package main

import (
	"errors"
	"fmt"
)

func Arithmetics(opt int, num1 int, num2 int) error {
	var result int
	if opt == 0 || num1 == 0 || num2 == 0 {
		return errors.New("Value cannot be 0!")
	}
	switch opt {
	case 1:
		fmt.Println("Addition")
		fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)
		result = num1 + num2
		fmt.Println(result)
	case 2:
		fmt.Println("Subtraction")
		fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)
		result = num1 - num2
		fmt.Println(result)
	case 3:
		fmt.Println("Multiplication")
		fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)
		result = num1 * num2
		fmt.Println(result)
	case 4:
		fmt.Println("Division")
		fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)
		result = num1 / num2
		fmt.Println(result)
	}
	return nil
}

func main() {
	err := Arithmetics(0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

}
