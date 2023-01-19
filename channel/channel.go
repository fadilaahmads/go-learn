package main

import (
	"fmt"
	"runtime"
)

func main() {
	// hannel digunakan untuk menghubungkan goroutine satu dengan goroutine lain.
	// Mekanisme yang dilakukan adalah serah-terima data lewat channel tersebut.
	// Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine,
	//  dan juga sebagai penerima di goroutine lainnya.
	// Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.

	// Set number logical processor cores
	runtime.GOMAXPROCS(2)
	// Create channel with make keyword
	var message = make(chan string)

	var sayHello = func(who string) {
		data := fmt.Sprintf("Hello, %s", who)
		message <- data
	}

	// Set goroutine
	go sayHello("Gunjo")
	go sayHello("Kenta")
	go sayHello("Yomi")

	var message1 = <-message
	fmt.Println(message1)
	var message2 = <-message
	fmt.Println(message2)
	var message3 = <-message
	fmt.Println(message3)

}
