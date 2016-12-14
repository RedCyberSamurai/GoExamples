package main

import (
	"fmt"
	"math/cmplx"
)

var (
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	i, j   int        = 1, 2
)

func main() {
	var c, python, java = true, false, "yes!"
	const f = "%T(%v)\n"
	/**
	 * T == Type
	 * v == variable
	 */

	fmt.Printf(f, i, z)
	fmt.Printf(f, z, i)
	fmt.Println(i, j, c, python, java, MaxInt, z)
}

// https://tour.golang.org/basics/10

/**
 * n ><= 0
 * int  int8  int16  int32  int64
 *
 * n >= 0
 * uint uint8 uint16 uint32 uint64 uintptr
 *
 * byte // alias for uint8
 *
 * rune // alias for int32
 *      // represents a Unicode code point
 *
 * float32 float64
 *
 * complex64 complex128
 *
 */
