// Note:
// Go Reader.Read
// The Reader.Read function reads data into a slice of bytes.

// func (b *Reader) Read(p []byte) (n int, err error)
// It returns the number of bytes read.

// In the next example we also use the hex package, which implements hexadecimal encoding and decoding.

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sid.jpg")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != nil {

			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		fmt.Printf("%s", hex.Dump(buf))
	}
}

// In the code example, we read an image and print it in hexadecimal format.

// reader := bufio.NewReader(f)
// We create a reader with bufio.NewReader.

// buf := make([]byte, 256)
// We create a custom buffer of 256 bytes.

// for {
//     _, err := reader.Read(buf)
// ...

// We read the binary data in a for loop.

// fmt.Printf("%s", hex.Dump(buf))
// The Dump returns a string that contains a hex dump of the given data.

// $ go run main.go
// 00000000  ff d8 ff e0 00 10 4a 46  49 46 00 01 01 00 00 01  |......JFIF......|
// 00000010  00 01 00 00 ff e1 00 2f  45 78 69 66 00 00 49 49  |......./Exif..II|
// 00000020  2a 00 08 00 00 00 01 00  0e 01 02 00 0d 00 00 00  |*...............|
// 00000030  1a 00 00 00 00 00 00 00  6b 69 6e 6f 70 6f 69 73  |........kinopois|
// 00000040  6b 2e 72 75 00 ff fe 00  3b 43 52 45 41 54 4f 52  |k.ru....;CREATOR|
// 00000050  3a 20 67 64 2d 6a 70 65  67 20 76 31 2e 30 20 28  |: gd-jpeg v1.0 (|
// 00000060  75 73 69 6e 67 20 49 4a  47 20 4a 50 45 47 20 76  |using IJG JPEG v|
// 00000070  38 30 29 2c 20 71 75 61  6c 69 74 79 20 3d 20 39  |80), quality = 9|
// 00000080  31 0a ff db 00 43 00 03  02 02 03 02 02 03 03 02  |1....C..........|
// 00000090  03 03 03 03 03 04 07 05  04 04 04 04 09 06 07 05  |................|
// 000000a0  07 0a 09 0b 0b 0a 09 0a  0a 0c 0d 11 0e 0c 0c 10  |................|
// 000000b0  0c 0a 0a 0e 14 0f 10 11  12 13 13 13 0b 0e 14 16  |................|
// ...
// In this tutorial, we have worked with the bufio package in Go.
