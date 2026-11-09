/*
SYNTAX
_____

switch expression {
case x,y:
   // code block if expression is evaluated to x or y
case v,w:
   // code block if expression is evaluated to v or w
case z:
...
default:
   // code block if expression is not found in any cases
}
*/

package main
import "fmt"

func main(){
	weather := "cloudy"

	switch weather {
	case "rainy", "cloudy":
		fmt.Println("bring a umbrella")
	case "sunny", "hot":
		fmt.Println("use a glasses")
	default:
		fmt.Println("sleep!!")
	}
	
	
}