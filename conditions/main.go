package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func newtonSqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 100000; i++ {
		z = z - ((math.Pow(z, 2) - x) / 2 * z)
	}

	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	fmt.Println("Printing square roots")
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println("Printing power")
	fmt.Println(pow(3, 2, 5))
	fmt.Println(pow(3, 2, 20))

	fmt.Println("Printing square roots @newton")
	fmt.Println(newtonSqrt(2))
}
