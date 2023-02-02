// Note:
// Go Writer.WriteRune
// The WriteRune writes a single rune.

// func (b *Writer) WriteRune(r rune) (size int, err error)
// It returns the number of bytes written and any error.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	runes := "🐜🐬🐄🐘🦂🐫🐑🦍🐯🐞"

	f, err := os.Create("runes.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	wr := bufio.NewWriter(f)

	for _, _rune := range runes {

		wr.WriteRune(_rune)
		wr.WriteRune('\n')
	}

	wr.Flush()

	fmt.Println("runes written")
}

// In the example, we read runes from a string and write them to a file; each one on a separate line.
