// 2. Error using Errorf() in Golang
// We can also handle Go errors using the Errorf() function. Unlike, New(), we can format the error message using Errorf().

// This function is present inside the fmt package, so we can directly use this if we have imported the fmt package.

// Let's see an example.

package main

import "fmt"

// func main() {

//   age := -14

//   // create an error using Efforf()
//   error := fmt.Errorf("%d is negative\nAge can't be negative", age)

//   if age < 0 {
//     fmt.Println(error)
//   } else {
//   fmt.Println("Age: %d", age);
//   }
// }
// Output

// -14 is negative
// Age can't be negative
// In the above example, we have used the Errorf() function to create a new formatted error.

// error := fmt.Errorf("%d is negative\nAge can't be negative", age)
// Here, you can see we have used the format specifier %d to use the value of age inside our error.

// --------------------------------------------------------------------------------------------------------------------------------

// Example: Error using Errorf():
// package main
// import "fmt"

func divide(num1, num2 int) error {

	// returns error
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}

	// returns the result of division
	return nil
}

func main() {

	err := divide(4, 0)

	// error found
	if err != nil {
		fmt.Printf("error: %s", err)

		// error not found
	} else {
		fmt.Println("Valid Division")
	}
}

// Output

// error: 4 / 0
// Cannot Divide a Number by zero
// In the above example, we have created a function named divide().

// func divide(num1, num2 int) error {...}
// The return of this function is an error, which means this function will return an error value.

// Inside the function, we have created a formatted error using the Errorf(). If the condition num2==0 is true, the function divide returns the error message inside the Errorf().

// However, if num2 is not 0, we are returning nil, which represents that there is no error.
