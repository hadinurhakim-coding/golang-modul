package main

import (
	"fmt"
	"time"
)

/*
   ========== GOROUTINE (Penjelasan Singkat) ==========

   Goroutine = "task ringan" yang dijalankan Go di belakang layar (concurrency).
   Dengan kata kunci "go" di depan pemanggilan fungsi, fungsi itu jalan di goroutine terpisah,
   sehingga program tidak menunggu selesai — bisa lanjut jalankan kode berikutnya.

   Analogi: kamu punya satu orang (main), lalu bilang "tolong cetak 1-5 di ruang lain" (go cetak()).
   Kamu tidak nunggu dia selesai; kamu langsung lanjut cetak A-E. Hasilnya bisa tercampur (interleaved).

   Penting: kalau main() selesai, semua goroutine ikut berhenti. Makanya sering pakai time.Sleep
   atau channel supaya main menunggu goroutine selesai.
*/

func cetakAngka() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func cetakHuruf() {
	for _, h := range "ABCDE" {
		fmt.Println("Huruf:", string(h))
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Tanpa "go": cetakAngka() jalan dulu sampai selesai, baru program lanjut (sequential).
	// Dengan "go": cetakAngka() jalan di goroutine terpisah; main langsung lanjut ke cetakHuruf().
	go cetakAngka()
	cetakHuruf() // jalan di main goroutine

	// Tanpa Sleep di bawah, main bisa selesai sebelum goroutine cetakAngka selesai, jadi output angka bisa tidak keluar semua.
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai.")
}
