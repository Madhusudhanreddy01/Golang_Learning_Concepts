// for index, val := range coll { }    |-------> Syntax
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, value := range nums { // "_ " is to ignore the index
		sum += value
	}
	fmt.Println("sum:", sum)

	fmt.Println("---------------------------------------------------------")

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	fmt.Println("---------------------------------------------------------")

	kvs := map[string]string{"1": "mango", "2": "apple", "3": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println("---------------------------------------------------------")

	for k := range kvs {
		fmt.Println("key:", k)
	}

	fmt.Println("---------------------------------------------------------")

	for i, c := range "Hi" {
		fmt.Println(i, c)
	}
}
