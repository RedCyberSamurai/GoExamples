package main

import (
	"fmt"
	"io"
	"strings"
	//"golang.org/x/tour/reader"
)

type CustomReader struct{}

/**
 * Overriding the Reader::Read() method
 */
func (r CustomReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil // bytestream length == 1
				  // error will not give an EOF, thats why it will continue infinitly
}

func main() {
	fmt.Println("--- Printing a simple bytestream ---")

	r := strings.NewReader("Some text to read in 8 byte strips.")
	b := make([]byte, 8)
	for {
		bytes, err := r.Read(b)
		fmt.Printf("NumBytes = %v Error = %v bytes in ASCII = %v\n", bytes, err, b)
		fmt.Printf("Current byte stream = %q\n", b[:bytes])

		if err == io.EOF { // detecting end of file
			break;
		}
	}

	fmt.Println("--- Manipulating bytestreams ---")
	// https://tour.golang.org/methods/22
}