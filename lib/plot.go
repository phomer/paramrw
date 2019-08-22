package lib

import "image"

type Plot struct {
	Pixels  *image.Paletted
	Xoffset int
	Yoffset int

	// Pen
	Color uint8
	Width int
}

func (plot *Plot) Set(x, y int, color uint8) {
	plot.Pixels.SetColorIndex(x+plot.Xoffset, y+plot.Yoffset, color)
}
