package main

import "fmt"

func main() {
//yang pertama: variabel tidak boleh dideklarasikan (:=) dua kali
//jika mau mengubah varibael cukup dengan (=)
	skor := 80
	pemenang := "MU"
	skor = 90

//di atas ada 2 variabel. jika salah satu variabel tidak dipakai, mana akan terjadi error.
//jika memang tidak mau memakai variabel tertentu, jalankan `[variabel] = _` -> variabel akan dianggap sudah digunakan
//
	fmt.Println("skor saya adalah", skor)
	fmt.Println("pemenangnya adalah", pemenang)

}
