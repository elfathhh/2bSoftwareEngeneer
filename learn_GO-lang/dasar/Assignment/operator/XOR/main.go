package main

import "fmt"

func main() {

  var x = 5

  fmt.Printf("x is %b \n", x)  // 101
  fmt.Printf("3 is %03b \n", 3) // 011

  x ^= 3

  fmt.Printf("x now is %03b \n", x) // 110
}