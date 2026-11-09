/*
SYNTAX
___
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}
*/

package main
import "fmt"

func main(){
	day := "Monday"

	switch day {
	case "Sunday":
		fmt.Println("Rest day")
	case "Monday":
		fmt.Println("go to village office")
	case "Tuesday":
		fmt.Println("go to TPQ")
	case "Wednesday":
		fmt.Println("turn to cook")
	case "Thursday":
		fmt.Println("rest -do something productive in posko")
	case "friday":
		fmt.Println("go to mosque")
	case "Saturday":
		fmt.Println("GO cycling")
	default:
		fmt.Println("GO HOME!!!")
	}
	
	
}