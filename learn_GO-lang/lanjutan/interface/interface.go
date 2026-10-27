// interface adalah kumpulan method signature yang bertindak
// sebagai "kontrak" perilaku. 
// Berbeda dengan bahasa lain yang menggunakan kata kunci khusus
// seperti implements, Go menggunakan implementasi implisit:
// sebuah struct otomatis dianggap mengimplementasikan interface jika
// ia memiliki semua method yang tercantum di dalam interface tersebut.

package main

import "fmt"

// 1. Definisikan Interface sebagai kontrak perilaku
type Pembayaran interface {
    Bayar(jumlah int)
}

// 2. Struct CreditCard
type CreditCard struct {
    NomorKartu string
}

func (c CreditCard) Bayar(jumlah int) {
    fmt.Printf("Membayar Rp%d menggunakan Credit Card (%s)\n", jumlah, c.NomorKartu)
}

// 3. Struct PayPal
type PayPal struct {
    Email string
}

func (p PayPal) Bayar(jumlah int) {
    fmt.Printf("Membayar Rp%d menggunakan PayPal (%s)\n", jumlah, p.Email)
}

// 4. Struct BankTransfer (Tipe data baru)
type BankTransfer struct {
    NamaBank string
}

func (b BankTransfer) Bayar(jumlah int) {
    fmt.Printf("Membayar Rp%d menggunakan Transfer Bank %s\n", jumlah, b.NamaBank)
}

// 5. Fungsi yang menerima Interface
func ProsesTransaksi(p Pembayaran, jumlah int) {
    p.Bayar(jumlah) // Memanggil method sesuai tipe aslinya secara otomatis
}

func main() {
    cc := CreditCard{NomorKartu: "1234-5678-9012"}
    pp := PayPal{Email: "user@email.com"}
    bt := BankTransfer{NamaBank: "BCA"}

    // Fungsi ProsesTransaksi tidak perlu diubah meskipun ada tipe pembayaran baru
    ProsesTransaksi(cc, 500000)
    ProsesTransaksi(pp, 250000)
    ProsesTransaksi(bt, 1000000)
}