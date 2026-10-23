// secara default, semua fungsi menerima data secara Pass by Value 
// (data dicopy/diduplikat).
// Namun dengan Pointer, 
// kita bisa mengirim data secara Pass by Reference (mengirim alamat memorinya).

package main

import "fmt"

// 1. Pass by Value (mengirim salinan nilai)
func ubahNilaiBiasa(n int) {
    n = 100 // Hanya mengubah nilai salinan di dalam fungsi
}

// 2. Pass by Reference (mengirim pointer/alamat memori)
func ubahNilaiPointer(n *int) {
    *n = 100 // Mengubah nilai asli di memori!
}

func main() {
    x := 10

    ubahNilaiBiasa(x)
    fmt.Println("menggunakan fungsi biasa:", x) // Nilai x TETAP 10

    ubahNilaiPointer(&x)
    fmt.Println("Setelah fungsi pointer:", x) // Nilai x BERUBAH jadi 100
}