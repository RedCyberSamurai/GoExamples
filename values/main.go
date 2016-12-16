package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 3)
}

func adder() func(int) int { // constructor for -> func(int) int
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	var n, k, f = 0, 1, 0
	return func() int {
		f = n + k
		n = k
		k = f
		return f
	}
}

func main() {
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Println("--- Printing values of different types ---")

	// q == display string surrounded by quotes
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("--- Using functions as values ---")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("--- Using functions as closures ---")
	inc, dec := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			inc(i),
			dec(-i),
		)
	}

	fmt.Println("--- Fibonacci ---")
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

}

// https://tour.golang.org/moretypes/26
