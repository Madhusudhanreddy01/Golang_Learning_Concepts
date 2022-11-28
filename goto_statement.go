package main

import (
	"fmt"
)

func main() {
	var input int
Loop:
	fmt.Print("You are not eligible to vote \n")
	fmt.Print("Enter your age: ")
	fmt.Scanln(&input)
	if input <= 17 {
		goto Loop
	} else {
		fmt.Print("You can vote ")
	}
}
