package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(num int) (x, y int) { // <<<
	x = num * 4 / 9
	y = num - x
	return // in the scope defined variables are retuned as integers
}

func main() {
	fmt.Println(add(5, 10))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
