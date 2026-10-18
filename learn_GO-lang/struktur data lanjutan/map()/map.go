//Jika Slice menyimpan data berdasarkan urutan indeks angka (0, 1, 2), 
// Map menyimpan data dalam bentuk pasangan Kunci dan Nilai (Key-Value).
// Map sangat cocok saat kita ingin mencari data berdasarkan kata kunci tertentu, bukan nomor urut.

package main

import "fmt"

func main(){
// Deklarasi Map: key berupa string, value berupa int
    hargaProduk := map[string]int{
        "Buku":  15000,
        "Pensil": 5000,
    }
// menambah ke map
	hargaProduk["Penghapus"] = 3000

	fmt.Println("Harga Produk:", hargaProduk["Buku"])
	fmt.Println("Semua Produk:", hargaProduk)
}