package main

import "fmt"

func main() {
	// Interface kosong atau empty interface yang dinotasikan dengan interface{} atau any (per go v1.18),
	// merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun.
	//  Tipe data dengan konsep ini biasa disebut dengan dynamic typing.

	var anything interface{}
	fmt.Println("Initial value", anything)
	// We can assign the "anything" variable with different data type
	anything = "this is string"
	fmt.Println(anything)
	anything = 8.56
	fmt.Println(anything)
	anything = [10]int{1, 23, 4, 55, 6, 87}
	fmt.Println(anything)

}
