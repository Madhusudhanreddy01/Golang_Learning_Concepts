// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Welcome to our Pizza App")
// 	fmt.Println("Please rate our Pizza between 1 and 5")

// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanks for Rating, ", input)

// 	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Added 1 to your rating:", numRating+1)
// 	}
// }

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Convert string to int")
	str_int := "10"
	val_int, _ := strconv.ParseInt(str_int, 10, 64)
	fmt.Printf("%d\n", val_int)

	fmt.Println("Convert string to int using atoi")
	val_atoi, _ := strconv.Atoi(str_int)
	fmt.Printf("%d\n", val_atoi)

	fmt.Println("Convert string to boolean")
	str_bool := "true"
	val_bool, _ := strconv.ParseBool(str_bool)
	fmt.Printf("%b\n", val_bool)

	fmt.Println("Convert string to float")
	str_float := "123.23"
	val_float, _ := strconv.ParseFloat(str_float, 64)
	fmt.Printf("%f\n", val_float)

	fmt.Println("Convert string to unsigned int")
	str_uint := "1234"
	val_uint, _ := strconv.ParseUint(str_uint, 10, 64)
	fmt.Printf("%d\n", val_uint)
}
