package main

import "fmt"

// Fungsi tanpa parameter dan tanpa nilai kembali (return)
func sapa() {
	fmt.Println("Hallo Hadi! Selamat datang di Go!")
}

// Fungsi dengan parameter dan tanpa nilai kembali (return)
func sapaNama(nama string) {
	fmt.Println("Hai", nama, "nagapain kamu!")
}

// Fungsi dengan parameter dan dengan nilai kembali (return)
func tambah(a int, b int) int {
	hasil := a + b
	return hasil
}

func main() {
	// Memanggil fungsi sapa tanpa parameter
	sapa()
	// Memanggil fungsi sapaNama dengan parameter "Hadi Nurhakim"
	sapaNama("Hadi Nurhakim")

	hasilTambah := tambah(5, 3)
	fmt.Println("5 + 3 =", hasilTambah)


	// Contoh fungsi dengan multiple return values
	luas, keliling := hitungPersegi(8)
	fmt.Println("Luas persegi dengan sisi 8 adalah:", luas)
	fmt.Println("Keliling persegi dengan sisi 8 adalah:", keliling)
}

	func hitungPersegi(sisi float64) (float64, float64) {
		luas := sisi * sisi
		keliling := 4 * sisi
		return luas, keliling
	}