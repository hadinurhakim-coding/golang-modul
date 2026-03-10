package main

import (
	"fmt"
	"time"
)

/*
   ========== CHANNEL (Penjelasan Singkat) ==========

   Channel = "pipa" untuk mengirim data dari satu goroutine ke goroutine lain (komunikasi aman).
   - make(chan T)    = channel unbuffered: kirim akan block sampai ada yang terima.
   - make(chan T, n) = channel buffered: bisa simpan n nilai tanpa block (sampai penuh).

   Operasi:
   - ch <- nilai  = mengirim nilai ke channel ch.
   - <-ch         = menerima nilai dari channel ch (bisa simpan: x := <-ch).
   - close(ch)    = menutup channel; penerima bisa tahu tidak ada data lagi.
*/

func kirimPesan(ch chan string, pesan string) {
	ch <- pesan // kirim ke channel (block sampai ada yang receive)
}

func main() {
	// Channel unbuffered: tipe string
	ch := make(chan string)

	go func() {
		ch <- "Hello dari goroutine!"
	}()

	// Main menunggu dan menerima dari channel
	terima := <-ch
	fmt.Println(terima) // Output: Hello dari goroutine!

	// Contoh buffered channel: kapasitas 2
	buf := make(chan int, 2)
	buf <- 1
	buf <- 2
	// buf <- 3  // kalau di-uncomment: deadlock (buffer penuh, tidak ada yang receive)
	fmt.Println(<-buf) // 1
	fmt.Println(<-buf) // 2

	// Contoh: goroutine mengirim beberapa nilai, main menerima
	ch2 := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(ch2)
	}()

	fmt.Println("Menerima dari ch2:")
	for v := range ch2 { // range channel: loop sampai channel di-close
		fmt.Println("  Terima:", v)
	}
	fmt.Println("Channel ditutup, selesai.")
}
