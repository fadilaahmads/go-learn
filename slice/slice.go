package main

import "fmt"

func main() {
	// Slice adalah reference elemen array.
	// Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
	// Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain
	// yang memiliki alamat memori yang sama.

	// Inisialisasi slice
	var veggies = []string{"Zuccini", "Spinach", "Kale", "Carrot"}
	fmt.Println(veggies)
	for _, vegetable := range veggies {
		fmt.Println(vegetable)
	}
}
