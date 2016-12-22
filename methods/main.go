package main

import (
	"math"
	"fmt"
)

type MyFloat float64
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/**
 * -- Pointer Receiver --
 * Same as a Receiver, but takes no return type.
 * Instead it modifies values from the existent instance of the given type.
 * Wether the given type is a pointer or not, the receiver pointer always grabs the pointer of the type.
 */
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

/**
 * -- Receiver --
 * Allows type MyFloat:: to have the method ::Abs()
 * func (t Type) Method() ReturnType {...}
 * => Type.Method()
 */
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/**
 * ScaleFunc(t Type, float64) => throws an error
 *         ---^---
 */
func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}


func main() {
	fmt.Println("Print a value from a receiver")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	fmt.Println("Print a value from a receiver modified by a pointer receiver")
	v := Vertex{3, 4}
	v.Scale(10) // {30, 40}
	fmt.Println(v.Abs())

	fmt.Println("Print a value from a function with pointer parameter")
	p := Vertex{3, 4}
	q := &Vertex{5, 6}
	ScaleFunc(&p, 8)
	ScaleFunc(q, 4)
	fmt.Println(p, q)
}

// https://tour.golang.org/methods/5