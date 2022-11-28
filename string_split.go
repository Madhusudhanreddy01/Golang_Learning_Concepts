// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "I,love,my,country"
// 	var arr []string = strings.Split(str, ",")
// 	fmt.Println(len(arr))
// 	for i, v := range arr {
// 		fmt.Println("Index : ", i, "value : ", v)
// 	}
// }

//----------------------------------------------------------------------

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("x,y,z", ","))
	fmt.Printf("%q\n", strings.Split(" John and Jack and Johnny and Jinn ", "and"))
	fmt.Printf("%q\n", strings.Split(" abc ", ""))
	fmt.Printf("%q\n", strings.Split("", "Hello"))
}
