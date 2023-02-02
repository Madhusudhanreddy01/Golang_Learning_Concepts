// Note:
// Reading file line by line with Scanner
// In the next example, we read a file line by line with a Scanner.

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("words.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// The example reads a small file containing words on each line.

// scanner := bufio.NewScanner(f)
// A new scanner is created with bufio.NewScanner.

// for scanner.Scan() {
//     fmt.Println(scanner.Text())
// }
// The Scan function advances the Scanner to the next token, which will then be available through the Bytes or Text method. By default, the function advances by lines.

//-----------------------------------------------------------------------------------------------------------

// Read words from a string
// In the following example, we read words from a string using Scanner.

package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {

	words := []string{}

	data := "A foggy mountain.\nAn old falcon.\nA wise man."

	sc := bufio.NewScanner(strings.NewReader(data))

	sc.Split(bufio.ScanWords)

	n := 0

	for sc.Scan() {
		words = append(words, sc.Text())
		n++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# of words: %d\n", n)

	for _, word := range words {

		fmt.Println(word)
	}
}

// The strings.NewReader returns a new reader from a string.

// sc.Split(bufio.ScanWords)
// We tell the scanner to scan by words using Split.
