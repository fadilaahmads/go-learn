package main

import "fmt"

type Animal struct {
	species string
	age     int
}

type Dog struct {
	animal Animal
	name   string
}

func (d Dog) Bark() string {
	bark := "Woof!"
	return bark
}
func (a Animal) getAge() int {
	age := a.age
	return age
}

func main() {
	// Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya).
	// Method bisa diakses lewat variabel objek.
	// Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct
	// hingga level private (level akses nantinya akan dibahas lebih detail pada chapter selanjutnya).
	// Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

	var dog1 = Dog{animal: Animal{"husky", 8}, name: "Buff"}
	fmt.Println(dog1.Bark())
	fmt.Println(dog1.animal.getAge())

}
