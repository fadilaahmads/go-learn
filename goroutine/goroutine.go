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
	// Specify the number of logical processor cores
	runtime.GOMAXPROCS(2)
	go Loop1(10, "With goroutine")
	Loop1(10, "Without goroutine")
	var input string
	fmt.Scanln("%s", &input)
}
