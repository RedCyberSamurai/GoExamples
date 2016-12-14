package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world") // executing last
	defer fmt.Println("1")
	defer fmt.Println("3")
	defer fmt.Println("2") // executing first
	// after all other executions below vvv

	/**
	 * Note: defer is a <Stack>
	 */

	fmt.Println("hello")
}
