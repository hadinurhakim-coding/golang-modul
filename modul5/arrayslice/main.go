package main

import "fmt"

func main() {
	// Deklarasi dan inisialisasi variabel
	var warna [4]string
	warna[0] = "merah"
	warna[1] = "kuning"
	warna[2] = "hijau"
	warna[3] = "biru"

	// Menampilkan nilai variabel
	fmt.Println("Array warna:", warna)
	fmt.Println("Warna ke 2:", warna[2])
}