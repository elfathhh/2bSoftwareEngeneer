package main

import "fmt"

func main() {

  var x = 5

  fmt.Printf("x is %b \n", x)  // 101

  x >>= 3	

  fmt.Printf("x now is %03b \n", x) // 000
}