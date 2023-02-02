// Custom Errors in Golang:
// In Go, we can create custom errors by implementing an error interface in a struct.

// error Interface:

// type error interface {
//   Error() string
// }
// Here, the Error() method returns an error message in string form if there is an error. Otherwise, it returns nil.

// Now, to create a custom error, we have to implement the Error() method on a Go struct.

// Let's see an example,

package main

import "fmt"

type DivisionByZero struct {
	message string
}

// define Error() method on the struct
func (z DivisionByZero) Error() string {
	return "Number Cannot Be Divided by Zero"
}

func divide(n1 int, n2 int) (int, error) {

	if n2 == 0 {
		return 0, &DivisionByZero{}
	} else {
		return n1 / n2, nil
	}
}

func main() {

	number1 := 15
	// change the value of number2 to get different result
	number2 := 0

	result, err := divide(number1, number2)

	// check if error occur or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
}

// Output

// Number Cannot Be Divided by Zero
// In the above example, we are implementing the Error() method of the error interface on the DivisionByZero struct.

// func (z DivisionByZero) Error() string {
//   return "Number Cannot Be Divided by Zero"
// }
// Here,

// z DivisionByZero - an instance of the DivisionByZero struct
// string - return type of the method
// "Number Cannot Be Divided by Zero" - error message
// We have then created a divide() method that takes two parameters and returns the result and an error.

// func divide(n1 int, n2 int) (int, error) {...}
// The function returns 0 and &DivisionByZero{}, if the value of n2 is 0. Here, &DivisionByZero{} is an instance of the struct. To learn more, visit Go Pointers to Struct.

// Inside the main() function, if the returned error type is not nil, we print the error message.

// Note that we are not calling the Error() method from anywhere in the program, but we can access its return value using the struct instance.
