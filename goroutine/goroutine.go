package main

import (
	"fmt"
	"runtime"
)

func Loop1(num int, message string) {
	for i := 0; i < num; i++ {
		fmt.Println(i, ":", message)
	}
}

func main() {
	// Goroutine mirip dengan thread, tapi sebenarnya bukan.
	// Sebuah native thread bisa berisikan sangat banyak goroutine.
	// Mungkin lebih pas kalau goroutine disebut sebagai mini thread.
	// Goroutine sangat ringan, hanya dibutuhkan sekitar 2kB memori saja untuk satu buah goroutine.
	// Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.

	// Specify the number of logical processor cores
	runtime.GOMAXPROCS(2)
	go Loop1(10, "With goroutine")
	Loop1(10, "Without goroutine")
	var input string
	fmt.Scanln("%s", &input)
}
