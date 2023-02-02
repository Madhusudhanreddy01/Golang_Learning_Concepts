// In programming, there will be situations where our program won't run properly and generate an error.
// For example,

package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		result := 20 / i
		fmt.Println(result)
	}
}

// When we run this code, we will get an error called integer divide by zero.

// During the first iteration, the value of i is 0, so the code result := 20 / i is trying to divide a number by 0.

// In this state, the program stops its execution known as a Go error.
// Here, "integer divide by zero" is an error message returned by the compiler.

// Note: The error we get in our previous example is a built-in error.
// In Go, we can also create our own errors for efficient programming.

// ------------------------------------------------------------------------------------------------------------------------

// Golang Error Handling:
// When an error occurs, the execution of a program stops completely with a built-in error message. In Go, we . Hence, it is important to handle those exceptions in our program.

// Unlike other programming languages, we don't use try/catch to handle errors in Go. We can handle errors using:

// 1. New() Function
// 2. Errof() Function
