package main

import (
	"fmt"
)

// Deklarasi struct
type Person struct {
	Nama         string
	Umur         int
	Tinggi       float64
	IsSudahNikah bool
	Hobi         []string
}

func main() {
	// Membuat objek dari struct
	var hadi Person // Deklarasi variabel dengan tipe struct Person
	hadi.Nama = "hadi" // Mengisi field Nama dengan nilai "hadi"
	hadi.Umur = 23 // Mengisi field Umur dengan nilai 23
	hadi.Tinggi = 167.5 // Mengisi field Tinggi dengan nilai 167.5
	hadi.IsSudahNikah = true // Mengisi field IsSudahNikah dengan nilai true
	hadi.Hobi = []string{"Coding", "main game", "nonton film", "naik gunung"} // Mengisi field Hobi dengan nilai ["Coding", "main game", "nonton film", "naik gunung"]

	fmt.Println("Hadi:", hadi) // Output: Hadi: {hadi 23 167.5 true [Coding main game nonton film naik gunung]}


	data := struct { // Deklarasi struct anonymous
		ID	int // Field ID dengan tipe int
		Title string // Field Title dengan tipe string
	}{ // Inisialisasi struct anonymous dengan nilai
		ID: 101, // Mengisi field ID dengan nilai 101
		Title: "Belajar Go", // Mengisi field Title dengan nilai "Belajar Go"
	}
	fmt.Println("Data anonymous:", data) // Output: Data anonymous: {101 Belajar Go}

	// Output: hadi berumur 23 tahun dan hobinya: [Coding main game nonton film naik gunung]
	fmt.Printf("%s berumur %d tahun dan hobinya: %v\n", hadi.Nama, hadi.Umur, hadi.Hobi)
}