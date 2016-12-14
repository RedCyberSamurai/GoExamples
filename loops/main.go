package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Sum integers by loop ---")
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}

	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	fmt.Println("--- For loop by range / foreach loop ---")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Println("2 ^", i, "=", v)
	}
}
