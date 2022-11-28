// const identifier [type] = value    |-----> Syntax

// const b string = "abc"     |------>Explicit Typing Example:

// const b = "abc"     |------>Implicit Typing Example:

package main

import "fmt"

func main() {
	const HEIGHT int = 10
	const WIDTH int = 20
	var area int
	area = HEIGHT * WIDTH
	fmt.Printf("value of area : %d", area)
}
