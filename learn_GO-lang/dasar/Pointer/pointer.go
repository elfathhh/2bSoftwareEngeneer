// adalah variabel yang menyimpan alamat memori dari variabel lain,
// bukan menyimpan nilainya langsung.

// Di Go, ada 2 simbol penting untuk Pointer:
// & (Ampersand): Digunakan untuk mengambil alamat memori suatu variabel.
// * (Asterisk): Digunakan untuk mengakses atau mengubah nilai yang ada di alamat memori tersebut.

package main

import "fmt"

func main(){
	angka := 10
	var ptr *int = &angka //ptr menyimpan alamat memori dari variabel angka

	fmt.Println("Nilai angka:", angka) // Output: 10
	fmt.Println("Alamat memori angka:", &angka) // Output: Alamat memori dari angka
	fmt.Println("Nilai via pointer:", *ptr) // Output: 10

	//mengubah nilai angka menggunakan pointer
	*ptr = 20
	fmt.Println("Nilai angka setelah diubah:", angka) // nilai 'angka' sudah di ubah ,melalui pointer '*"" 
}
