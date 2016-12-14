package main

import (
	"fmt"
)

func main() {
	var (
		i int
		f float64
		b bool
		s string
	)

	// q == display string surrounded by quotes
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
