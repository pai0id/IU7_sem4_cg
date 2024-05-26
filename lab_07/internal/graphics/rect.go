package graphics

import (
	"image/color"
)

type Rect struct {
	P1, P2 FPoint
}

func DrawRect(pixels *SafePixels, rect Rect, c color.Color) {
	p1, p2 := ftoiPoint(rect.P1), ftoiPoint(rect.P2)

	var xSt, xEd int
	if p1.X < p2.X {
		xSt, xEd = p1.X, p2.X
	} else {
		xSt, xEd = p2.X, p1.X
	}
	for x := xSt; x <= xEd; x++ {
		setPixel(pixels, IPoint{x, p1.Y}, c)
		setPixel(pixels, IPoint{x, p2.Y}, c)
	}

	var ySt, yEd int
	if p1.Y < p2.Y {
		ySt, yEd = p1.Y, p2.Y
	} else {
		ySt, yEd = p2.Y, p1.Y
	}
	for y := ySt; y <= yEd; y++ {
		setPixel(pixels, IPoint{p1.X, y}, c)
		setPixel(pixels, IPoint{p2.X, y}, c)
	}
}
