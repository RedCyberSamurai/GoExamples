package main

import (
	"fmt"
)

type Callable interface {
	CallName()
}

/**
 * An interface is a set of methods.
 * Receivers cannot be defined
 */
type Actor struct {
	name string
}

/**
 * This will meet the condition to use the interface "Callable"
 */
func (a Actor) CallName() {
	fmt.Println("Calling Actor:", a.name)
}


func main() {
	fmt.Println("--- Using interface ---")
	var c Callable
	c = Actor{"John Doe"} // other types who do not assign the same methods, will throw an error

	c.CallName()

	fmt.Println("--- Using interface returning a nil receiver ---")
	var c2 Callable
	var a *Actor // nil == nullptr (?)
	c2 = a
	describe(c2)
	// a.CallName() <= would cause a runtime error

	fmt.Println("--- Using interface to handle any type ---")
	var i interface{}
	describeAny(i)

	i = 42
	describeAny(i)

	i = "Some String"
	describeAny(i)
}

func describe(c Callable) {
	fmt.Printf("(%v, %T)\n", c, c)
}

func describeAny(c interface{}) {
	fmt.Printf("(%v, %T)\n", c, c)
}