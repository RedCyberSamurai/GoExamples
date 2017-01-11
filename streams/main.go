package main

import (
	"fmt"
	"io"
	"strings"
	//"golang.org/x/tour/reader"
	"golang.org/x/tour/reader"
	"os"
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


type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	r := rune(b)

	if r < 'n' && r >= 'a' || r < 'N' && r >= 'A' {
		b += 13
	} else if r >= 'n' && r <= 'z' || r >= 'N' && r <= 'Z' {
		b -= 13
	}

	return b
}

func (rot rot13Reader) Read(b []byte) (out int, err error) {

	out, err = rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		b[i] = rot13(b[i])
	}

	return
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
	reader.Validate(CustomReader{})
	// https://tour.golang.org/methods/22

	fmt.Println("--- Override Reader for rot13 encryption ---")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rotReader := rot13Reader{s}
	io.Copy(os.Stdout, &rotReader)
}