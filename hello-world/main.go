package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Something printable.")
	fmt.Println("Hello World!")

	fmt.Println("The time is", time.Now())

	rand.Seed(time.Now().Unix())
	fmt.Println("Random Number:", rand.Intn(10))

	fmt.Println("This is a decimal: ", math.Sqrt(7))
	fmt.Println("This is PI: ", math.Pi)

	// run package (on sublime) with:
	// :replay
	//
	// build executable with:
	// go build
}
