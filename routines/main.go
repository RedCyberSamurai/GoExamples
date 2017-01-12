package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("--- Parallel execution of a function ---")
	go say("world")
	say("hello")


}
