package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("--- Printing an array ---")
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println("--- Printing an array of primes ---")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/**
	 * Slice Syntax:
	 * [|x:|y] contains range x to y-1  [x:y]
	 * [n|:] contains range n+1 to max  [n:]
	 * [:n|] contains range min to n    [:n]
	 *
	 * Note: a slice can be extended by using a lower/higher value than the given min/max
	 */

	fmt.Println("--- Printing a slice of an array ---")
	var s []int = primes[1:5] // <<
	fmt.Println(s)

	fmt.Println("--- Printing a slice literals ---")
	lit := []struct {
		i int
		b bool
	}{
		{2, true}, {3, false}, {5, true}, {7, true}, {11, false}, {13, true},
	}

	fmt.Println(lit)

	fmt.Println("--- Printing length and capacity of a slice ---")
	fmt.Println("The array", primes)
	fmt.Println("Slice", s)
	fmt.Println("Length", len(s))
	fmt.Println("Capacity", cap(s))

	fmt.Println("--- Printing nilable slice ---")
	var nilable []int
	fmt.Println(nilable, len(nilable), cap(nilable))
	if nilable == nil {
		fmt.Println("Slice is nil!")
	}

	fmt.Println("--- Printing declared dynamic array ---")
	dyn := make([]int, 5)
	fmt.Println(dyn, len(dyn), cap(dyn))

	fmt.Println("--- Printing slices of slices ---")
	board := [][]string{ // tic tac toe
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// taking turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("--- Printing appending slices ---")
	var toAppend []int
	fmt.Println("before", toAppend, len(toAppend), cap(toAppend))

	toAppend = append(toAppend, 1, 2, 3)

	fmt.Println("after", toAppend, len(toAppend), cap(toAppend))
}

// https://tour.golang.org/moretypes/8
