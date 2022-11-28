package main

import "fmt"

func main() {
	x := 5
	y := 10
	// fmt.Println("-----------------------")

	swap(&x, &y)

	fmt.Println(x, y)
}

func swap(x *int, y *int) {

	*x, *y = *y, *x
	*x = 20
	*y = 40
	fmt.Println(*x)
	fmt.Println(*y)
	fmt.Println("-----------------------")

}
