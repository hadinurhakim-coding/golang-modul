package calculator

/*
   ========== PACKAGE: calculator ==========

   Tujuan package ini:
   - Menunjukkan cara membuat package sendiri.
   - Menunjukkan perbedaan identifier Exported vs Unexported.

   Aturan export di Go sangat sederhana:
   - Nama yang diawali huruf BESAR → bisa diakses dari package lain (exported).
   - Nama yang diawali huruf kecil → hanya bisa dipakai di dalam package ini (unexported).
*/

// Add adalah fungsi exported (bisa dipanggil dari package lain).
func Add(a, b int) int {
	return a + b
}

// Sub adalah fungsi exported (bisa dipanggil dari package lain).
func Sub(a, b int) int {
	return a - b
}

// multiply adalah fungsi unexported (hanya dipakai internal package).
func multiply(a, b int) int {
	return a * b
}

// Mul adalah wrapper exported yang memanfaatkan fungsi unexported.
func Mul(a, b int) int {
	return multiply(a, b)
}

