package main

import "fmt"

func main() {

	// Percabangan if-else
	umur := 20

	if umur >= 18 {
		fmt.Println("Dewasa!")
	} else {
		fmt.Println("Anak kecil")
	}

	// Perulangan for
	for i := 1; i <= 5; i++ { 
		fmt.Println("Loop ke-", i)
	}

	// Switch case
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Jangan lupa ini hari senin.")
	default:
		fmt.Println("Sekarang bukan hari senin.")
	}
}