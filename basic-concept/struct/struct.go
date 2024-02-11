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
type College struct {
	student []Student
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
	fmt.Println("Mahasiswa1:", mahasiswa1)

	var mhs2 = Person{name: "Stevi", age: 20}
	var mahasiswa2 = Student{id: 187388490, Person: mhs2, subject: availableSubject}
	fmt.Println("Mahasiswa2:", mahasiswa2)

	allStudents := []Student{mahasiswa1, mahasiswa2}
	var collegeStudents = College{
		student: allStudents,
	}
	for _, student := range collegeStudents.student {
		fmt.Println(student)
	}

}
