package main

// https://go.dev/tour/methods/25

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	color  color.Model
	bounds image.Rectangle
}

func (i Image) ColorModel() color.Model {
	return i.color
}

func (i Image) Bounds() image.Rectangle {
	return i.bounds
}

func (i Image) At(x, y int) color.Color {

	r := uint8(x*1 + y*2)
	g := uint8((x*1 + y*2) / 2)
	b := uint8(55)
	a := uint8(255)

	return color.RGBA{r, g, b, a}
}

func main() {
	m := Image{
		bounds: image.Rect(0, 0, 1920, 1080),
		color:  color.RGBAModel,
	}
	pic.ShowImage(m)
}
