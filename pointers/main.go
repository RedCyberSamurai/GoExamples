package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	fmt.Println("Accessing i and print from pointer")
	p := &i
	fmt.Println(*p)

	fmt.Println("Accessing j and modify")
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
