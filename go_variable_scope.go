// // Program to illustrate local variables

// package main

// import "fmt"

// func addNumbers() {

// 	// local variables
// 	var sum int

// 	sum = 5 + 9

// }

// func main() {

// 	addNumbers()

// 	// cannot access sum out of its local scope
// 	fmt.Println("Sum is", sum)

// }

// -------------------------------------------------------------------------

// Program to illustrate global variable

package main

import "fmt"

// declare global variable before main function
var sum int

func addNumbers() {

	// local variable
	sum = 9 + 5
}

func main() {

	addNumbers()

	// can access sum
	fmt.Println("Sum is", sum)

}
