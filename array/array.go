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

	var angka = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(angka[0])

	// Multidimension array
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	for i := 0; i < len(buah); i++ {
		fmt.Println(buah[i])
	}

	for i, num := range angka {
		fmt.Printf("index ke: %d, nilainya: %d\n", i, num)
	}
	// for _, num := range angka {
	// 	fmt.Printf("index ke: %d, nilainya: %d\n", i, num)
	// }

	var warna = make([]string, 5)
	warna[0] = "merah"
	warna[1] = "jingga"
	warna[2] = "kuning"
	warna[3] = "hijau"
	warna[4] = "biru"

	for _, color := range warna {
		fmt.Println(color)
	}
}
