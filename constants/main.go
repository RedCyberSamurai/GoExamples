package main

import (
	"fmt"
)

const Pi = 3.14

const (
	Big   = 1 << 100  // 1 followed by 100 zeroes
	Small = Big >> 99 // shift 1 -> 99 places to the right == 2^1
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "世界"
	fmt.Println("Hello", World)

	fmt.Println("Happy", Pi, "Day")

	const Boolean = false
	fmt.Println("False is:", Boolean)

	fmt.Println("Big - Small prints")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
