package main

import "fmt"

// Deklarasi struct untuk merepresentasikan sebuah persegi panjang
type Reactangle struct {
	Panjang float64  // Field Panjang dengan tipe float64 untuk menyimpan panjang persegi panjang
	Lebar  float64 // Field Lebar dengan tipe float64 untuk menyimpan lebar persegi panjang
}


func (r *Reactangle) Luas() float64 { // Method Luas untuk menghitung luas persegi panjang
	return r.Panjang * r.Lebar // Mengembalikan hasil perkalian antara panjang dan lebar sebagai luas
}

func (r *Reactangle) Keliling() float64 { // Method Keliling untuk menghitung Keliling persegi panjang
	return 2 * (r.Panjang + r.Lebar) // Mengembalikan hasil penjumlahan panjang dan lebar, kemudian dikalikan 2 sebagai Keliling
}

func (r *Reactangle) Scale(factor float64) { // Method Scale untuk mengubah ukuran persegi panjang
	r.Panjang *= factor // Mengalikan panjang dengan faktor untuk mengubah ukuran
	r.Lebar *= factor // Mengalikan lebar dengan faktor untuk mengubah ukuran
}

func main() {

	kotak := Reactangle{Panjang: 10, Lebar: 5} // Membuat objek kotak dengan panjang 10 dan lebar 5

	fmt.Printf("Kotak: %+v\n", kotak) // Output: Kotak: {Panjang:10 Lebar:5}
	fmt.Println("Luas:", kotak.Luas()) // Output: Luas: 50
	fmt.Println("Keliling", kotak.Keliling()) // Output: Keliling: 30	

	kotak.Scale(2) // Mengubah ukuran kotak dengan faktor 2
	fmt.Printf("Setelah scale 2x: %+v\n", kotak) // Output: Setelah scale 2x: {Panjang:20 Lebar:10}
	fmt.Println("Luas setelah scale:", kotak.Luas()) // Output: Luas setelah scale: 200
	fmt.Println("Keliling setelah scale:", kotak.Keliling()) // Output: Keliling setelah scale: 60
}