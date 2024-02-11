package main

import "fmt"

// Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
// Dengan menerapkan konsep tersebut, kita bisa membuat fungsi di dalam fungsi,
// atau bahkan membuat fungsi yang mengembalikan fungsi.

func main() {
	var compare = func(num1, num2 int) string {
		var result string
		if num1 > num2 {
			result = fmt.Sprintf("Num1 with %d has bigger value than Num2 %d", num1, num2)
		} else if num1 < num2 {
			result = fmt.Sprintf("Num1 with %d has smaller value than Num2 %d", num1, num2)
		} else {
			result = fmt.Sprintf("Num1 with %d has same value compared to Num2 %d", num1, num2)
		}
		return result
	}
	var angka1 int
	var angka2 int
	fmt.Scanf("%d", &angka1)
	fmt.Scanf("%d", &angka2)
	input := compare(angka1, angka2)
	fmt.Println(input)

}
