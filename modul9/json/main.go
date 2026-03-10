package main

import (
	"encoding/json"
	"fmt"
)

/*
   ========== JSON (Penjelasan Singkat) ==========

   encoding/json dipakai untuk:
   - Marshal   = struct → JSON string (byte slice).
   - Unmarshal = JSON string (byte slice) → struct.

   Struct tag `json:"nama_field"` dipakai agar nama field di JSON bisa beda dengan nama field di Go,
   atau agar field diabaikan dengan `json:"-"`.
*/

type User struct {
	Nama  string `json:"nama"`
	Umur  int    `json:"umur"`
	Email string `json:"email,omitempty"` // omitempty: tidak muncul di JSON kalau kosong
}

func main() {
	// Struct → JSON (Marshal)
	u := User{Nama: "Hadi", Umur: 23, Email: "hadi@contoh.com"}
	bytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error Marshal:", err)
		return
	}
	fmt.Println("JSON (Marshal):", string(bytes))

	// JSON → Struct (Unmarshal)
	jsonStr := `{"nama":"Dina","umur":22,"email":""}`
	var u2 User
	err = json.Unmarshal([]byte(jsonStr), &u2)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}
	fmt.Printf("Struct (Unmarshal): %+v\n", u2)

	// Marshal dengan indent (lebih enak dibaca)
	bytesIndent, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("JSON indent:")
	fmt.Println(string(bytesIndent))
}
