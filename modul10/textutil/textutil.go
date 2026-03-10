package textutil

import (
	"strings"
	"unicode"
)

/*
   ========== PACKAGE: textutil ==========

   Tujuan package ini:
   - Contoh package helper kecil untuk string.
   - Contoh penggunaan package standard library: strings, unicode.
*/

// Title membuat huruf pertama tiap kata menjadi kapital (contoh sederhana).
// Catatan: ini implementasi sederhana untuk belajar; untuk kebutuhan kompleks bisa pakai library lain.
func Title(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	kata := strings.Fields(s)
	for i := range kata {
		r := []rune(kata[i])
		if len(r) == 0 {
			continue
		}
		r[0] = unicode.ToUpper(r[0])
		kata[i] = string(r)
	}
	return strings.Join(kata, " ")
}

// HasPrefix adalah wrapper kecil untuk menunjukkan pemanggilan fungsi dari stdlib.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

