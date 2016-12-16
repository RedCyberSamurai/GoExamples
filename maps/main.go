package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var (
	map1 map[string]Vertex
	map2 map[string]Vertex
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)

	for i := 0; i < len(f); i++ {
		m[f[i]]++
	}

	return m
}

func main() {
	fmt.Println("Printing values from a map")

	map1 = make(map[string]Vertex)
	map1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(map1["Bell Labs"])

	fmt.Println("Printing values from a map using literals")
	map2 = map[string]Vertex{ // <<<
		"Bell Labs": Vertex{ // "Vertex" can be left out, because of the given return type
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(map2)

	fmt.Println("Delete a value from a key [Bell Labs]")
	delete(map2, "Bell Labs")
	fmt.Println("Value: ", map2["Bell Labs"])

	fmt.Println("Check if key of map exists")
	value, isPresent := map2["Bell Labs"]
	fmt.Println("[Bell Labs] Value:", value, "Is present:", isPresent)
	value, isPresent = map2["Google"]
	fmt.Println("[Google] Value:", value, "Is present:", isPresent)

	fmt.Println("Word count with help of map")
	wc.Test(WordCount)
}
