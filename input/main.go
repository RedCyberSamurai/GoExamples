package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	fmt.Println("--- Reading console input ---")
	// for some reason fmt.Scan() does not like spaces...
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	if !scanner.Scan() {
		println("Could not scan input")
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		println(err)
		return
	}

	fmt.Println("This is an integer: ", num)

	if !scanner.Scan() {
		println("Could not scan input")
	}

	t := scanner.Text()
	flo, err := strconv.ParseFloat(t, len(t))

	if err != nil {
		println(err)
		return
	}
	fmt.Printf("This is a float: %T %v", flo, flo)

}
