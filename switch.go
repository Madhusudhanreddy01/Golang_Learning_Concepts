// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print("Enter Number: ")
// 	var input int
// 	fmt.Scanln(&input)
// 	switch input {
// 	case 10:
// 		fmt.Print("the value is 10")
// 	case 20:
// 		fmt.Print("the value is 20")
// 	case 30:
// 		fmt.Print("the value is 30")
// 	case 40:
// 		fmt.Print("the value is 40")
// 	default:
// 		fmt.Print("Entered something wrong ")
// 	}
// }

// ______________________________________________________________________________________________

// package main

// import "fmt"

// func main() {
// 	k := 30
// 	switch k {
// 	case 10:
// 		fmt.Println("was <= 10")
// 		fallthrough
// 	case 20:
// 		fmt.Println("was <= 20")
// 		fallthrough
// 	case 30:
// 		fmt.Println("was <= 30")
// 		fallthrough
// 	case 40:
// 		fmt.Println("was <= 40")
// 		fallthrough
// 	case 50:
// 		fmt.Println("was <= 50")
// 		fallthrough
// 	case 60:
// 		fmt.Println("was <= 60")
// 		fallthrough
// 	default:
// 		fmt.Println("default case")
// 	}
// }

// ________________________________________________________________

// // Golang program to show the
// // uses of fallthrough keyword
package main

// Here "fmt" is formatted IO which
// is same as C’s printf and scanf.
import "fmt"

// Main function
func main() {
	day := "Tue"

	// Use switch on the day variable.
	switch {
	case day == "Mon":
		fmt.Println("Monday")
		fallthrough
	case day == "Tue":
		fmt.Println("Tuesday")
		fallthrough
	case day == "Wed":
		fmt.Println("Wednesday")
	}
}
