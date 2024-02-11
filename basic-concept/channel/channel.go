package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	// =======
	// channel
	// =======
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

	// ====================
	// Channel as parameter
	// ====================
	// Variabel channel bisa di-pass ke fungsi lain sebagai parameter.
	// Cukup tambahkan keyword chan pada deklarasi parameter agar operasi pass channel variabel bisa dilakukan.
	// Passing data bertipe channel lewat parameter sifatnya pass by reference,
	// yang di transferkan adalah pointer datanya, bukan nilai datanya.
	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			message <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(message)
	}
}
