package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type RGBA image.RGBA

var cyan = color.RGBA{100, 200, 200, 0xff}

func (img *RGBA) fillInnerSquare(x0, y0, length, factor int) {
	//Fill center
	for x := x0 + length/3; x < x0+2*(length/3); x++ {
		for y := y0 + length/3; y < y0+2*(length/3); y++ {
			(*image.RGBA)(img).Set(x, y, cyan)
		}
	}

	if factor == 1 || length == 3 {
		return
	}

	// Fill neighbours if possible
	img.fillInnerSquare(x0, y0, length/3, factor-1)
	img.fillInnerSquare(x0+length/3, y0, length/3, factor-1)
	img.fillInnerSquare(x0+2*(length/3), y0, length/3, factor-1)

	img.fillInnerSquare(x0, y0+length/3, length/3, factor-1)
	img.fillInnerSquare(x0+2*(length/3), y0+length/3, length/3, factor-1)

	img.fillInnerSquare(x0, y0+2*(length/3), length/3, factor-1)
	img.fillInnerSquare(x0+length/3, y0+2*(length/3), length/3, factor-1)
	img.fillInnerSquare(x0+2*(length/3), y0+2*(length/3), length/3, factor-1)
}

func main() {
	pixelDepth := 8
	width := (int)(math.Pow(3, float64(pixelDepth)))

	factor := 6

	// Image build
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, width}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel
	(*RGBA)(img).fillInnerSquare(0, 0, width, factor)

	// Encode as PNG
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
