package main

import "fmt"

// Declaring struct
type Person struct {
	name string
	age  int
}

type Student struct {
	id int
	Person
	subject []string
}

func main() {
	// Struct initialization
	var mahasiswa1 = Student{}
	mahasiswa1.name = "Diki"
	mahasiswa1.age = 19
	mahasiswa1.id = 1852305884
	availableSubject := []string{
		"IT fundamental",
		"Networking",
		"Cybersecurity",
		"Web Programming",
		"Desktop Programming",
		"Embeded System",
	}
	mahasiswa1.subject = append(mahasiswa1.subject, availableSubject...)

	fmt.Println(mahasiswa1)

}
