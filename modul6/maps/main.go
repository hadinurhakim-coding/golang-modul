package main

import "fmt"

func main() {
    // Map adalah tipe data yang menyimpan pasangan key-value
	umur := make(map[string]int)

	// Menambahkan pasangan key-value ke map
	umur["Hadi"] = 23
	umur["Dina"] = 22
	umur["Rizky"] = 24

	// Menampilkan nilai map
	fmt.Println("Map umur:", umur) // Output: Map umur: map[Dina:22 Hadi:23 Rizky:24]

	// Mengakses nilai map dengan key
	fmt.Println("Umur Hadi:", umur["Hadi"]) // Output: Umur Hadi: 23
	
	if val, ok := umur["Hadi"]; ok { // Cek apakah key "Hadi" ada di map
		fmt.Println("Umur Hadi:", val) // Jika ada, tampilkan nilainya
	} else { // Jika tidak ada, tampilkan pesan
		fmt.Println("Hadi ga ada di map") // Output: Hadi ga ada di map
	}


	nilai := map[string]float64{ // Deklarasi dan inisialisasi map dengan literal
		"Matematika": 85.5, // Nilai untuk mata pelajaran Matematika
		"Bahasa Indonesia": 90.0, // Nilai untuk mata pelajaran Bahasa Indonesia
	}
	fmt.Println("Nilai Hadi:", nilai) // Output: Nilai Hadi: map[Bahasa Indonesia:90 Matematika:85.5]
}