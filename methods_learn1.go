// Consider the following Go program. How can you call print() ?

package main

import "fmt"

type usd float64

func (u usd) print() {
	fmt.Printf("USD value is: %v\n", u)
}

func main() {
	var balance usd = 50.5
	// Answer for Above question
	balance.print()

}
