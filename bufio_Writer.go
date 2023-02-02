// Note:
// A new writer is created with bufio.NewWriter or bufio.NewWriterSize.

// func NewWriter(w io.Writer) *Writer
// func NewWriterSize(w io.Writer, size int) *Writer

// The NewWriter function returns a new Writer whose buffer has the default size.
// The NewWriterSize returns a new Writer whose buffer has at least the specified size.

// Go Writer.WriteString
// The WriteString writes a string to the buffer.

// func (b *Writer) WriteString(s string) (int, error)
// It returns the number of bytes written.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := []string{"an old falcon", "misty mountains", "a wise man", "a rainy morning"}

	f, err := os.Create("words.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	wr := bufio.NewWriter(f)
	for _, line := range data {
		wr.WriteString(line + "\n")
	}
	wr.Flush()
	fmt.Println("data written")
}

// The example writes a few strings into a file with Writer.WriteString.

// wr := bufio.NewWriter(f)
// A new writer is created. The default buffer size is 4KB.

// for _, line := range data {

//     wr.WriteString(line + "\n")
// }
// In a for loop, we write the data to the buffer.

// wr.Flush()
// Since our data is smaller than the default 4KB buffer size, we have to call Flush for the data to be actually written to the file.
