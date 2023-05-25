package main

import (
	"image/color"
	"math"

	"github.com/tidwall/pinhole"
)

type pinHole pinhole.Pinhole

var cyan = color.RGBA{100, 200, 200, 0xff}

func (p *pinHole) fillInnerSquare(x0, y0, z0, length float64, factor int) {
	//Fill each face and center

	(*pinhole.Pinhole)(p).DrawCube(
		x0, y0+length/3, z0+length/3,
		x0+length/3, y0+2*(length/3), z0+2*(length/3),
	)
	(*pinhole.Pinhole)(p).DrawCube(
		x0+2*(length/3), y0+length/3, z0+length/3,
		x0+length, y0+2*(length/3), z0+2*(length/3),
	)

	// center
	(*pinhole.Pinhole)(p).DrawCube(
		x0+length/3, y0+length/3, z0+length/3,
		x0+2*(length/3), y0+2*(length/3), z0+2*(length/3),
	)

	(*pinhole.Pinhole)(p).DrawCube(
		x0+length/3, y0, z0+length/3,
		x0+2*(length/3), y0+2*(length/3), z0+2*(length/3),
	)
	(*pinhole.Pinhole)(p).DrawCube(
		x0+length/3, y0+2*(length/3), z0+length/3,
		x0+2*(length/3), y0+length, z0+2*(length/3),
	)

	(*pinhole.Pinhole)(p).DrawCube(
		x0+length/3, y0+length/3, z0,
		x0+2*(length/3), y0+2*(length/3), z0+2*(length/3),
	)
	(*pinhole.Pinhole)(p).DrawCube(
		x0+length/3, y0+length/3, z0+2*(length/3),
		x0+2*(length/3), y0+2*(length/3), z0+length,
	)

	if factor == 1 {
		return
	}

	// Recursion
	p.fillInnerSquare(x0, y0, z0, length/3, factor-1)
	p.fillInnerSquare(x0+length/3, y0, z0, length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0, z0, length/3, factor-1)

	p.fillInnerSquare(x0, y0+length/3, z0, length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0+length/3, z0, length/3, factor-1)

	p.fillInnerSquare(x0, y0+2*(length/3), z0, length/3, factor-1)
	p.fillInnerSquare(x0+length/3, y0+2*(length/3), z0, length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0+2*(length/3), z0, length/3, factor-1)
	//
	p.fillInnerSquare(x0, y0, z0+length/3, length/3, factor-1)
	p.fillInnerSquare(x0+length/3, y0, z0+length/3, length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0, z0+length/3, length/3, factor-1)

	p.fillInnerSquare(x0, y0+length/3, z0+length/3, length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0+length/3, z0+length/3, length/3, factor-1)

	p.fillInnerSquare(x0, y0+2*(length/3), z0+length/3, length/3, factor-1)
	p.fillInnerSquare(x0+length/3, y0+2*(length/3), z0+length/3, length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0+2*(length/3), z0+length/3, length/3, factor-1)
	//
	p.fillInnerSquare(x0, y0, z0+2*(length/3), length/3, factor-1)
	p.fillInnerSquare(x0+length/3, y0, z0+2*(length/3), length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0, z0+2*(length/3), length/3, factor-1)

	p.fillInnerSquare(x0, y0+length/3, z0+2*(length/3), length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0+length/3, z0+2*(length/3), length/3, factor-1)

	p.fillInnerSquare(x0, y0+2*(length/3), z0+2*(length/3), length/3, factor-1)
	p.fillInnerSquare(x0+length/3, y0+2*(length/3), z0+2*(length/3), length/3, factor-1)
	p.fillInnerSquare(x0+2*(length/3), y0+2*(length/3), z0+2*(length/3), length/3, factor-1)
}

func main() {
	factor := 1

	p := pinhole.New()

	p.Begin()

	p.DrawCube(-0.3, -0.3, -0.3, 0.3, 0.3, 0.3)

	(*pinHole)(p).fillInnerSquare(-0.3, -0.3, -0.3, 0.6, factor)

	p.Scale(1.5, 1.5, 1.5)
	p.Rotate(math.Pi/3, math.Pi/6, 0)
	p.End()

	p.SavePNG("cube.png", 500, 500, &pinhole.ImageOptions{
		Scale:     1,
		BGColor:   color.White,
		LineWidth: 0.25,
	})
}
