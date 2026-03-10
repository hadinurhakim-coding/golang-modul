package main

import (
	"fmt"
	"math"
)

// Interface mendefinisikan perilaku (method) yang harus dipenuhi
type Bentuk interface {
	Luas() float64
	Keliling() float64
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.JariJari
}

// CetakLuasKeliling menerima interface Bentuk (bisa Persegi atau Lingkaran)
func CetakLuasKeliling(b Bentuk) {
	fmt.Printf("Luas: %.2f, Keliling: %.2f\n", b.Luas(), b.Keliling())
}

func main() {
	persegi := Persegi{Sisi: 5}
	lingkaran := Lingkaran{JariJari: 3}

	fmt.Println("Persegi sisi 5:")
	CetakLuasKeliling(persegi) // Output: Luas: 25.00, Keliling: 20.00

	fmt.Println("Lingkaran jari-jari 3:")
	CetakLuasKeliling(lingkaran) // Output: Luas: 28.27..., Keliling: 18.84...

	// Slice of interface: kumpulkan berbagai bentuk
	bentukBentuk := []Bentuk{persegi, lingkaran}
	fmt.Println("Daftar bentuk:")
	for i, b := range bentukBentuk {
		fmt.Printf("  %d - Luas: %.2f\n", i+1, b.Luas())
	}
}
