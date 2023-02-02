// 1. Go Error using New() Function:
// In Go, we can use the New() function to handle an error.
// This function is defined inside the errors package and allows us to create our own error message.

// Let's see an example,

package main

// import the errors package
import (
	"errors"
	"fmt"
)

// func main() {

// 	message := "Hello"

// 	// create error using New() function
// 	myError := errors.New("WRONG MESSAGE")

// 	if message != "Programiz" {
// 		fmt.Println(myError)
// 	}

// }

// Output

// WRONG MESSAGE
// In the above example, we have created an error using the errors.New() function.

// myError := errors.New("WRONG MESSAGE")
// Here, the "WRONG MESSAGE" inside the New() function is the custom error message. It prints when the message variable does not match with the given string "PROGRAMIZ".

// -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// Example: Error using New()
// package main

// // import the errors package
// import (
//   "errors"
//   "fmt"
// )

// function that checks if name is Programiz
func checkName(name string) error {

	// create a new error
	newError := errors.New("Invalid Name")

	// return the error if name is not Programiz
	if name != "Programiz" {
		return newError
	}

	// return nil if there is no error
	return nil
}

func main() {

	name := "Hello"

	// call the function
	err := checkName(name)

	// check if the err is nil or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Name")
	}

}

// Output

// Invalid Name
// In the above example, we have created a function named checkName()

// checkName(name string) error {...}
// The return of this function is an error, which means this function will return a value of error type

// Inside the function, we have created an error using the New() function. Here, if the name is not Programiz, we are returning the newly created error message.

// However, if the name is Programiz, we are returning nil (represents no error).

// Inside the main() function, we have called the checkName() function using

// err := checkName(name)

// Here, the returned error will be assigned to err. We then checked if the value in err is nil or not.
