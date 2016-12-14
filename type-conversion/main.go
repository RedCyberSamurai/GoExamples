package main

import (
	"fmt"
)

func main() {
	/**
	 * Same as:
	 *
	 * i := 42
	 * f := float64(i)
	 * u := uint(f)
	 *
	 */
	var (
		i int     = 42
		f float64 = float64(i)
		u uint    = uint(f)
	)

	const p = "%T(%v) %T(%v) %T(%v)\n"
	fmt.Printf(p, i, i, f, f, u, u)

	v1 := 42
	v2 := 42.0
	v3 := "string"
	v4 := 0.5i
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v3 is of type %T\n", v3)
	fmt.Printf("v4 is of type %T\n", v4)
}
