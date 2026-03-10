package main

import "fmt"

/*
   ========== DEFER (Penjelasan Singkat) ==========

   defer = "tunda eksekusi" sampai fungsi yang melingkupinya selesai (return atau panic).
   Berguna untuk: menutup file, koneksi DB, unlock mutex — supaya tetap jalan bahkan kalau ada return/error.

   Urutan: LIFO (Last In, First Out). Defer yang terakhir dipanggil, pertama dijalankan saat fungsi selesai.
*/

func main() {
	// Defer akan jalan saat main() selesai (setelah semua kode di main selesai)
	defer fmt.Println("Ini cetak terakhir (defer 1)")
	defer fmt.Println("Ini cetak kedua dari belakang (defer 2)")
	fmt.Println("Ini cetak pertama")

	// Output urutan:
	// Ini cetak pertama
	// Ini cetak kedua dari belakang (defer 2)
	// Ini cetak terakhir (defer 1)

	// Contoh praktis: defer untuk "cleanup"
	fmt.Println("--- Contoh defer di fungsi ---")
	cetakDenganDefer()
}

func cetakDenganDefer() {
	defer fmt.Println("  [cleanup] fungsi selesai, ini selalu jalan")
	fmt.Println("  Mulai kerja...")
	fmt.Println("  Selesai kerja.")
	// Di sini (sebelum return), semua defer dijalankan dari yang terakhir didefer
}
