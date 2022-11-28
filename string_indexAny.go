// Go program to illustrate how to find
// the index value of the given string
package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Creating and initializing the strings
	str1 := "Welcome to the online portal of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Finding the index value
	// of the given strings
	// Using  IndexAny() function
	res1 := strings.IndexAny(str1, "G")
	res2 := strings.IndexAny(str2, "do")
	res3 := strings.IndexAny(str3, "qxa")
	res4 := strings.IndexAny("GeeksforGeeks, geeks", "uywq")

	// Displaying the result
	fmt.Println("\nIndex values:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)

}
