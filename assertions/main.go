package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Printing type assertions ---")
	var i interface{} = "to assert"

	s := i.(string)
	fmt.Println(s)

	s, assert := i.(string)
	fmt.Println(s, assert)

	f, assert := i.(float64)
	fmt.Println(f, assert)

	// Because the value is no float, you cannot assign it to a variable.
	// Instead it throws a panic error.
	f = i.(float64)
	fmt.Println(f)
}