package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	str := input
	a := FirstNonRepeating(str)
	fmt.Print(a)
}
func FirstNonRepeating(str string) string {
	for _, c := range str {
		if strings.Count(str, string(c)) < 2 {
			return string(c)
		}
	}
	return ""
}
