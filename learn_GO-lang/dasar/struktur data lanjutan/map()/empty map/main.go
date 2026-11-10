/*
There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.

SYNTAX
___
var a map[KeyType]ValueType
*/

// This example shows the difference between 
// declaring an empty map using with the make()function and without it.

package main
import "fmt"

func main() {
  var a = make(map[string]string)
  var b map[string]string

  fmt.Println(a == nil)
  fmt.Println(b == nil)
}