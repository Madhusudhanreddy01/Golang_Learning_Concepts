package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23 //:= walrus operator

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}

//-------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := 10
// 	changeX(&x)
// 	fmt.Println(x)
// }
// func changeX(x *int) {
// 	*x = 0
// }
