//**SLICE & ARRAY**
//sama aja sebenarnya. hanya array di go itu punya ukuran tetap - tidak bisa diubah
//berbeda dengan slice yang lebih fleksibel. 
// itu juga yang menyebabkan slice lebih sering dipakai

package main

import "fmt"

func main(){
// 1. Array
	var array = [3]int{1, 2, 3}
//=> dimulai dengan mendeklarasikan panjang array dan menentukan tipe data yg dipakai.

// 2. Slice
	angkaSlice := []int{1, 2, 3}
//=> dimulai dengan mendeklarasikan slice tanpa panjang (di dalam square brackets) tetap dan menentukan tipe data yg dipakai.

// 2.1 menambah slice
	angkaSlice = append(angkaSlice, 4)

	fmt.Println("ini adalah array:", array)
	fmt.Println("ini adalah slice:", angkaSlice)
}

//apakah ada parameter untuk mengubah isi slice selain append? 
// bisa, tapi harus menggunakan index. 
// misal angkaSlice[0] = 10. maka angkaSlice akan menjadi [10, 2, 3, 4]