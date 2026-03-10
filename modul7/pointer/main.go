package main

import "fmt"

/*
   ========== APA ITU POINTER? (Bahasa Awam) ==========

   Bayangkan variabel itu seperti "kotak" yang punya dua hal:
   1. ISI kotak   = nilai yang kita simpan (misal: angka 25)
   2. ALAMAT kotak = lokasi kotak itu di gudang (di memori komputer)

   Biasanya kita hanya peduli ISI-nya. Tapi kadang kita perlu bilang ke komputer:
   "Jangan kasih saya kopi isinya—kasih saya ALAMAT kotaknya saja."
   Itu yang disebut POINTER: kita pegang "alamat", bukan "isi"-nya langsung.

   Perumpamaan sehari-hari:
   - Variabel biasa = kamu punya RUMAH (dan isi rumah itu datanya).
   - Pointer       = kamu punya KARTU ALAMAT yang menunjuk ke rumah itu.
                     Lewat kartu alamat, orang bisa datangi rumah itu dan mengubah isinya.
                     Jadi perubahan terasa di "rumah asli", bukan di fotokopi.

   Kenapa pakai pointer?
   - Supaya fungsi bisa MENGUBAH nilai asli (bukan cuma kopi). Tanpa pointer,
     yang diubah cuma "tembusan"-nya; yang asli tetap.
   - Supaya tidak perlu menyalin data besar—cukup kirim "alamat"-nya saja (lebih hemat).

   Simbol penting:
   - & (ampersand) = "alamat dari..."  → &umur = alamat di mana umur disimpan.
   - * (asterisk)  = "isi yang ada di alamat ini..." → *ptr = nilai yang ditunjuk ptr.
   - * di tipe ( *int ) = "pointer yang menunjuk ke nilai int".
*/

// Address adalah struct (wadah) untuk data alamat. Harus didefinisikan dulu sebelum dipakai.
type Address struct {
	Kota     string
	Provinsi string
	Negara   string
}

func main() {

	// --- Variabel biasa: kita simpan nilai 25 di "kotak" bernama umur ---
	umur := 25

	// ptr = "kartu alamat" yang menunjuk ke kotak umur. &umur = "di mana kotak umur berada?"
	var ptr *int = &umur
	fmt.Println("Nilai umur:", umur)             // Isi kotak: 25
	fmt.Println("Alamat memori umur:", &umur)    // Alamat kotak (angka hex = lokasi di memori)
	fmt.Println("Pointer ptr menunjuk ke:", ptr) // ptr berisi alamat yang sama
	fmt.Println("Nilai yang ditunjuk ptr (*ptr):", *ptr) // Buka alamat itu, ambil isinya = 25 (dereference)

	// Mengubah isi "kotak asli" lewat pointer: kita ke alamat itu, lalu ganti isinya jadi 30
	*ptr = 30
	fmt.Println("Setelah *ptr = 30, umur menjadi:", umur) // Kotak asli ikut berubah: 30

	// Pointer yang belum menunjuk ke mana-mana = nil (seperti kartu alamat kosong)
	var ptrKosong *int
	fmt.Println("Pointer kosong:", ptrKosong) // Output: <nil>

	// Contoh praktis: fungsi double ingin mengubah angka ASLI, bukan kopinya.
	// Kita kirim "alamat" angka (&angka), sehingga perubahan di dalam fungsi terasa di sini.
	angka := 10
	fmt.Println("Sebelum double:", angka) // 10
	double(&angka)                        // Kirim alamat kotak "angka"
	fmt.Println("Sesudah double:", angka) // Kotak asli berubah jadi 20

	alamat1 := Address{Kota: "tasik", Provinsi: "garut", Negara: "bandung"}
	alamat2 := &alamat1

	alamat2.Kota = "jakarta"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

// double menerima "alamat" sebuah kotak (n *int). Isi kotak di alamat itu kita kali 2.
func double(n *int) {
	*n = *n * 2 // *n = baca isi di alamat n, lalu tulis kembali hasil kali 2 ke alamat yang sama
}
