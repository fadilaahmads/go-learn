package main

import "fmt"

// Looping seperti for biasa
func Loop1(a int, b int) {
	end := b
	for start := a; start <= end; start++ {
		fmt.Println(start)
	}
}

// Looping seperti while
func Loop2(x int, z int) {
	for x < z {
		fmt.Println(x)
		x += 1
	}
}

// Continue akan menskip ke iterasi berikutnya
// Break akan menghentikan looping
func Loop3(x int, z int) {
	end := z

	for start := x; start < end; start++ {
		if start%2 == 0 {
			continue
		}
		if start > 8 {
			break
		}
		fmt.Println(start)
	}
}

func main() {
	Loop1(1, 3)
	fmt.Println("-----")
	Loop2(1, 5)
	fmt.Println("-----")
	Loop3(0, 10)
}
