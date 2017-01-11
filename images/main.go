package main

import (
	"fmt"
	"image"
)

func main() {

	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println("--- Print rectangle coordinates ---")
	fmt.Println(m.Bounds())
	fmt.Println("--- Print image pixel rgba values ---")
	fmt.Println(m.At(0,0).RGBA())
}