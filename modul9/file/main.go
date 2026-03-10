package main

import (
	"fmt"
	"os"
)

/*
   ========== BACA/TULIS FILE (Penjelasan Singkat) ==========

   os.ReadFile(namaFile)  = baca seluruh isi file → []byte, error.
   os.WriteFile(namaFile, data, permission) = tulis data ke file (overwrite). Permission contoh: 0644.

   Untuk file besar atau streaming, pakai os.Open / os.Create dan bufio.Scanner atau io.Copy.
*/

func main() {
	// Menulis file
	namaFile := "contoh.txt"
	isi := "Halo, ini isi file dari modul9/file.\nBaris kedua."
	err := os.WriteFile(namaFile, []byte(isi), 0644)
	if err != nil {
		fmt.Println("Error tulis file:", err)
		return
	}
	fmt.Println("File ditulis:", namaFile)

	// Membaca file
	data, err := os.ReadFile(namaFile)
	if err != nil {
		fmt.Println("Error baca file:", err)
		return
	}
	fmt.Println("Isi file:")
	fmt.Println(string(data))

	// Opsional: hapus file contoh setelah selesai (biar tidak numpuk)
	// os.Remove(namaFile)
}
