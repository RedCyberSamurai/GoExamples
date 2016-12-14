package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	pV = &Vertex{1, 2}
)

func main() {
	fmt.Println("Printing a struct")
	fmt.Println(Vertex{1, 2})

	fmt.Println("Modify new declared struct")
	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v.X)

	fmt.Println("Modify struct by pointer")
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println("Printing structs defined by specific literals")
	fmt.Println(v1, v2, v3, pV)
}
