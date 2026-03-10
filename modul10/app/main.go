package main

import (
	"fmt"

	// Import package buatan sendiri menggunakan module path dari go.mod.
	// Format umumnya:
	//   <module-name>/<folder-package>
	//
	// Karena go.mod kita:
	//   module github.com/hadinurhakim-coding/golang-modul
	// maka import package modul10 adalah:
	//   github.com/hadinurhakim-coding/golang-modul/modul10/<nama-package>
	"github.com/hadinurhakim-coding/golang-modul/modul10/calculator"
	"github.com/hadinurhakim-coding/golang-modul/modul10/textutil"
)

func main() {
	fmt.Println("=== Modul10: Package & Import ===")

	// Pakai package calculator
	fmt.Println("calculator.Add(10, 5) =", calculator.Add(10, 5))
	fmt.Println("calculator.Sub(10, 5) =", calculator.Sub(10, 5))
	fmt.Println("calculator.Mul(10, 5) =", calculator.Mul(10, 5))

	// Pakai package textutil
	kalimat := "belajar golang itu seru"
	fmt.Println("textutil.Title(...) =", textutil.Title(kalimat))
	fmt.Println("textutil.HasPrefix(\"golang\", \"go\") =", textutil.HasPrefix("golang", "go"))

	// Catatan penting:
	// - Kamu tidak bisa memanggil calculator.multiply(...) dari sini karena itu unexported (huruf kecil).
	// - Kalau ingin bisa dipakai dari luar package, namanya harus diawali huruf besar.
}

