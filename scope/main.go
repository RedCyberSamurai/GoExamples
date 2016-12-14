package main

import (
	"fmt"
)

func main() {
	/**
	 * := is only defining a variable in the current scope
	 */

	i := 42
	fmt.Println("You can see me [i]", i)
	cmd()
}

func cmd() {
	fmt.Println("You cant see me [::cmd]", i) // will throw an error
}
