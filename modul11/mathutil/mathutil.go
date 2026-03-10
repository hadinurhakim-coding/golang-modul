package mathutil

import "errors"

/*
   ========== MODUL11: TESTING ==========

   Package ini sengaja dibuat kecil untuk latihan unit test dan benchmark.
*/

// Sum menjumlahkan semua angka dalam slice.
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Divide membagi a dengan b dan mengembalikan error jika b == 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return a / b, nil
}

