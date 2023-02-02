// Note:
// A new reader is created with bufio.NewReader or bufio.NewReaderSize.

// func NewReader(rd io.Reader) *Reader
// func NewReaderSize(rd io.Reader, size int) *Reader

// The NewReader function returns a new Reader whose buffer has the default size.
// The NewReaderSize returns a new Reader whose buffer has at least the specified size.

// Go Reader.ReadString
// The ReadString reads until the first occurrence of the given delimiter in the input.

// func (b *Reader) ReadString(delim byte) (string, error)
// It returns a string containing the data up to and including the delimiter.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter your name: ")
	r := bufio.NewReader(os.Stdin)
	name, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s!\n", strings.TrimSpace(name))
}

// With ReadString function, we read an input from the user and produce a message to the console.

// r := bufio.NewReader(os.Stdin)
// We create a new reader from the standard input.

// name, err := r.ReadString('\n')
// A string input is read form the user.
