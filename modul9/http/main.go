package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
   ========== HTTP CLIENT (Penjelasan Singkat) ==========

   http.Get(url) mengembalikan *http.Response dan error.
   Response punya Body (io.ReadCloser) yang harus di-Close setelah selesai dibaca.
   Baca isi Body dengan io.ReadAll(resp.Body).
*/

func main() {
	// Contoh: request GET ke API publik (JSON placeholder)
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status tidak OK:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error baca body:", err)
		return
	}

	fmt.Println("Status:", resp.Status)
	panjang := 200
	if len(body) < panjang {
		panjang = len(body)
	}
	fmt.Println("Body (potongan):", string(body)[:panjang], "...")
}
