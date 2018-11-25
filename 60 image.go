package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
	V             byte
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width, m.Height)
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{m.V, m.V, 255, 255}
}

func main() {
	m := Image{100, 200, 127}
	pic.ShowImage(m)
}
