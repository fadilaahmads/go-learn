package main

import "fmt"

func ChangeWitoutPointer(ori int, change int) int {
	ori = change
	return ori
}

func ChangeWithPointer(ori *int, change int) int {
	ori = &change
	return *ori
}

func main() {
	// Pointer adalah reference atau alamat memori.
	// Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
	// Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4,
	// maka yang dimaksud pointer adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri.

	// Inisialisasi pointer
	var varBiasa int = 1980
	var varPointer *int = &varBiasa

	fmt.Println("Nilai varBiasa:", varBiasa)
	fmt.Println("Alamat memori varBiasa:", varPointer)

	// dereference variabel
	fmt.Println("Dereference varPointer:", *varPointer)

	// Pointer as parameter
	num := 9
	fmt.Println("Num before:", num)
	noPointer := ChangeWitoutPointer(num, 5)
	fmt.Println("Num after:", noPointer)
	withPointer := ChangeWithPointer(&num, 5)
	fmt.Println("Num after:", withPointer)
}
