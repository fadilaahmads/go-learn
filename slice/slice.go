package main

import "fmt"

func main() {
	// Slice adalah reference elemen array.
	// Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
	// Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain
	// yang memiliki alamat memori yang sama.

	// Inisialisasi slice
	var veggies = []string{
		"Zuccini",
		"Spinach",
		"Kale",
		"Carrot",
		"Aubergine",
		"Daikon",
		"Birds Eye Chili",
		"Bean Sprouts",
	}
	fmt.Println(veggies)
	for _, vegetable := range veggies {
		fmt.Println(vegetable)
	}

	// pick veggies on index 0 to 4
	fmt.Println("veggies on index 0 to 4", veggies[0:6])
	// slice from slice
	sVeggies := veggies[3:6]
	fmt.Println("Slice from slice veggies:", sVeggies)

	// Panjang slice
	sliceSize := len(veggies)
	fmt.Println("Slice Size:", sliceSize)

}
