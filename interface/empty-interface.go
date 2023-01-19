package main

import "fmt"

func main() {
	// Interface kosong atau empty interface yang dinotasikan dengan interface{} atau any (per go v1.18),
	// merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun.
	//  Tipe data dengan konsep ini biasa disebut dengan dynamic typing.

	// ===============
	// Interface usage
	// ===============
	var anything interface{}
	fmt.Println("Initial value", anything)
	// We can assign the "anything" variable with different data type
	anything = "this is string"
	fmt.Println("Anything string", anything)
	anything = 8.56
	fmt.Println("Anything float", anything)
	anything = [10]int{1, 23, 4, 55, 6, 87}
	fmt.Println("Anything Array", anything)

	// =============================
	// Empty interface Variable cast
	// =============================
	// Variabel bertipe interface{} bisa ditampilkan ke layar sebagai string dengan memanfaatkan fungsi print, seperti fmt.Println().
	// Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah nilai asli, melainkan bentuk string dari nilai aslinya.
	// Hal ini penting diketahui, karena untuk melakukan operasi yang membutuhkan nilai asli pada variabel yang bertipe interface{},
	// diperlukan casting ke tipe aslinya.

	var secret interface{}
	secret = 2
	// fmt.Println(secret * 2) <<--- this line error because: [invalid operation: secret * 2 (mismatched types interface{} and int)]
	// We had to convert from intreface{} data type to int
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is", number)
	secret = []string{"Adobo", "Burgir", "Croissant"}
	for _, food := range secret.([]string) {
		fmt.Println(food)
	}

	// =======================================
	// Slice, map, and interface{} combination
	// =======================================
	// tipe data tersebut bisa menjadi alternatif tipe slice struct.
	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

}
