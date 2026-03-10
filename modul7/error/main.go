package main

import (
	"errors"
	"fmt"
)

// Pembagi mengembalikan hasil bagi dan error jika pembagi 0
func Pembagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return a / b, nil
}

// UmurValid memvalidasi umur, mengembalikan error jika tidak valid
func UmurValid(umur int) error {
	if umur < 0 || umur > 120 {
		return fmt.Errorf("umur %d tidak valid (harus 0-120)", umur)
	}
	return nil
}

func main() {
	// Contoh 1: penanganan error dari Pembagi
	hasil, err := Pembagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("10 / 2 =", hasil) // Output: 10 / 2 = 5

	hasil, err = Pembagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: pembagi tidak boleh nol
	} else {
		fmt.Println("Hasil:", hasil)
	}

	// Contoh 2: validasi dengan fmt.Errorf
	if err := UmurValid(25); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Umur 25 valid")
	}

	if err := UmurValid(150); err != nil {
		fmt.Println("Error:", err) // Output: Error: umur 150 tidak valid (harus 0-120)
	}
}
