package main

import (
	"fmt"
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	Width, Height int
	Color uint8
}

// the dimension of the image
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

// the color type of the image
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// the color on a coordinate of the image
func (i Image) At(x, y int) color.Color {
	return color.RGBA{R:i.Color + uint8(x), G:i.Color + uint8(y), B:255, A:255}
}

func main() {

	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println("--- Print rectangle coordinates ---")
	fmt.Println(m.Bounds())
	fmt.Println("--- Print image pixel rgba values ---")
	fmt.Println(m.At(0,0).RGBA())

	fmt.Println("--- Generating Image by definition ---")
	m2 := Image{100, 100, 185}
	pic.ShowImage(m2)
}