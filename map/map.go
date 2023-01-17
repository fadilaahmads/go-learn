package main

import "fmt"

func main() {
	// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair.
	// Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya.
	// Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
	// Map mirip array/slice tapi aksesnya bukan dari indeks, pakai key-value pair

	// Inisialisasi map
	// map[tipe key]tipe value
	var foods map[string]int
	// insialisasi nilai dalam map
	foods = map[string]int{} // biar gak nil

	foods["Apple Pie"] = 20
	foods["Burgir"] = 14
	foods["Cake"] = 19
	foods["Dimsum"] = 15
	foods["Edamame"] = 5
	foods["Falafel"] = 8
	foods["Grape Juice"] = 3

	fmt.Println(foods["Burgir"])
	for food, price := range foods {
		fmt.Printf("Menu: %s, price: %d\n", food, price)
	}

	// Cara inisialisasi map 2
	var drinks = map[string]int{
		"Apple Juice":  3,
		"Black Coffee": 6,
		"Cider":        2,
		"Dr. Peppers":  4,
		"Fizzy Soda":   5,
	}

	fmt.Println("Drinks", drinks)

}
