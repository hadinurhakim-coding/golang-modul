package main

import "fmt"

func main() {
	var nama string = "Hadi"
	umur := 25
	tinggi := 167.5
	isDeveloper := true

	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
	fmt.Println("Developer:", isDeveloper)

	total := umur + 5
	fmt.Println("5 tahun lagi umur saya adalah:", total)
}