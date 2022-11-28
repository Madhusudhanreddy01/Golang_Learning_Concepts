// prices := 200
// prices := 300
// prices := 400

// package main

// func main() {
// // var prices = [3]int{200, 300, 400}

// var length int

// fmt.Scan(&length)
// fmt.Println("Enter :", length)

// nums := []int{}

// for _, prices := range nums {
// 	fmt.Scan(&prices)
// 	nums = append(nums, prices)
// }
// fmt.Println(nums)

// for i := 0; i < length; i++ {
// 	fmt.Scan(&prices)
// 	nums = append(nums, prices)
// }
// fmt.Println(nums)

// prices[1] = 600
// fmt.Println(prices)

// }

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{1, 2, 3, 20}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 33, 34, 35, 36, 38)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}
