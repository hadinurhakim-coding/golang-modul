package main

import (
	"fmt"
	"math"
)

/*
   ========== APA ITU INTERFACE? (Penjelasan Detail) ==========

   Interface di Go = "KONTRAK PERILAKU". Kita tidak peduli "benda"-nya apa (struct apa),
   yang penting benda itu BISA melakukan hal-hal yang kita sebutkan di interface.

   Analogi:
   - Interface "Bentuk" bilang: "Siapa pun yang mau jadi Bentuk HARUS punya method Luas() dan Keliling()."
   - Persegi punya Luas() dan Keliling() → Persegi memenuhi kontrak → Persegi BISA dipakai sebagai Bentuk.
   - Lingkaran punya Luas() dan Keliling() → Lingkaran juga memenuhi kontrak → Lingkaran BISA dipakai sebagai Bentuk.
   - Jadi satu fungsi CetakLuasKeliling(b Bentuk) bisa menerima Persegi ATAU Lingkaran tanpa perlu dua fungsi terpisah.

   Keuntungan:
   1. Polimorfisme: satu fungsi untuk banyak tipe (Persegi, Lingkaran, Segitiga, dll).
   2. Kode lebih fleksibel: nanti tambah "Segitiga", cukup buat struct + method Luas/Keliling, langsung bisa dipakai.
   3. Di Go: implementasi interface IMPLICIT (tidak pakai kata kunci "implements"). Asal methodnya ada, otomatis memenuhi interface.
*/

// Bentuk adalah interface: mendefinisikan "kontrak" — tipe apa pun yang punya Luas() dan Keliling() float64 bisa dipakai sebagai Bentuk.
type Bentuk interface {
	Luas() float64
	Keliling() float64
}

// --- Implementasi pertama: Persegi ---
type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// --- Implementasi kedua: Lingkaran ---
type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.JariJari
}

// CetakLuasKeliling menerima parameter bertipe interface Bentuk.
// Artinya: bisa terima Persegi, Lingkaran, atau tipe lain yang punya Luas() dan Keliling().
// Di dalam fungsi kita hanya memanggil b.Luas() dan b.Keliling() — tidak peduli aslinya Persegi atau Lingkaran.
func CetakLuasKeliling(b Bentuk) {
	fmt.Printf("Luas: %.2f, Keliling: %.2f\n", b.Luas(), b.Keliling())
}

func main() {
	persegi := Persegi{Sisi: 5}
	lingkaran := Lingkaran{JariJari: 3}

	// Satu fungsi yang sama bisa dipakai untuk Persegi dan Lingkaran (polimorfisme).
	fmt.Println("Persegi sisi 5:")
	CetakLuasKeliling(persegi) // Output: Luas: 25.00, Keliling: 20.00

	fmt.Println("Lingkaran jari-jari 3:")
	CetakLuasKeliling(lingkaran) // Output: Luas: 28.27..., Keliling: 18.84...

	// Slice of interface: kita bisa mengumpulkan berbagai "bentuk" dalam satu slice.
	// Masing-masing elemen bisa Persegi atau Lingkaran; yang penting semuanya punya Luas() dan Keliling().
	bentukBentuk := []Bentuk{persegi, lingkaran}
	fmt.Println("Daftar bentuk:")
	for i, b := range bentukBentuk {
		fmt.Printf("  %d - Luas: %.2f\n", i+1, b.Luas())
	}
}
