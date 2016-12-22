package main

import (
	"fmt"
	"time"
)

type CustomError struct {
	When time.Time
	What string
}

/**
 * Overrides the method from error::Error()
 */
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error occured at %v.\nError code: %s", e.When, e.What)
}

func main() {
	fmt.Println("--- Printing overrided error output ---")
	err := &CustomError{
		time.Now(),
		"Some error code 0x0000",
	}
	if err != nil {
		fmt.Println(err)
	}
}

// https://tour.golang.org/methods/20