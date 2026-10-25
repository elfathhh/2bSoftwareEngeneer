// Jika Slice dan Map menyimpan kumpulan data dengan tipe yang seragam,
// Struct digunakan untuk menggabungkan beberapa data dengan tipe data yang berbeda
// menjadi satu kesatuan objek. Ini sangat mirip dengan tabel di database!

package main

import "fmt"

// Definisi blueprint Struct
type Produk struct {
    ID    int
    Nama  string
    Harga int
}

func main() {
    // Membuat instance dari Struct
    p1 := Produk{
        ID:    1,
        Nama:  "Laptop",
        Harga: 12000000,
    }

    fmt.Println("Nama Produk:", p1.Nama) // Mengakses field menggunakan tanda titik (.)
}