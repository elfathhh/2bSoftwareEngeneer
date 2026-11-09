//The for loop can take up to three statements:
/*
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
*/
//EXPLANATION :
/*
statement1 Initializes the loop counter value.

statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

statement3 Increases the loop counter value.
*/

package main
import "fmt"

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}