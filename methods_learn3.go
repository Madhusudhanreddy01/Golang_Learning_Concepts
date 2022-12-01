// What will the value of balance  be after running the following program?

package main

import "fmt"

type usd float64

func (u *usd) zero() {
	*u = 0
}

func main() {
	var balance usd = 50.5
	balance.zero()
	fmt.Println(balance)
}

// The method changes the value the pointer points to.
