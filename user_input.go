package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a number:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error While Reading Input") //nil error
	}

	trimmedinput := strings.TrimSpace(input)         //"42\n(delimeter)"
	n, err := strconv.ParseInt(trimmedinput, 10, 32) //strings to integer
	if err != nil {
		log.Fatal("Input was not a number")
	}
	fmt.Println(n)
}
