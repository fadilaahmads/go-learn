package main

import "fmt"

func main() {
	// Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel
	// Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan,
	// menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan.

	// inisialisasi array
	var buah [6]string
	// isi array
	buah[0] = "apple"
	buah[1] = "semangka"
	buah[2] = "jeruk"
	buah[3] = "melon"
	buah[4] = "leci"
	buah[5] = "rambutan"

	fmt.Println(buah[3])
}
