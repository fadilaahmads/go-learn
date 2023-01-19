package main

import "math"

// Interface definition
type hitung interface {
	luas() float64
	keliling() float64
}

// Lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// Persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	// Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja),
	// yang dibungkus dengan nama tertentu.
	// Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil.
	// Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki definisi method
	// minimal sama dengan yang ada di interface-nya

}
