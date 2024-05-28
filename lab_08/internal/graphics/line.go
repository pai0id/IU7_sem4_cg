package graphics

import (
	"image/color"
	"math"
)

type Line struct {
	P1, P2 FPoint
}

func DrawLine(pixels *SafePixels, l Line, c color.Color) {
	p1, p2 := l.P1, l.P2

	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	dx := x2 - x1
	dy := y2 - y1

	steps := math.Max(math.Abs(dx), math.Abs(dy))

	xIncrement := dx / steps
	yIncrement := dy / steps

	x := x1
	y := y1

	for i := 0; float64(i) <= steps; i++ {
		setPixel(pixels, IPoint{X: round(x), Y: round(y)}, c)
		x += xIncrement
		y += yIncrement
	}
}

func drawLine(pixels *SafePixels, l Line, c color.Color) {
	p1, p2 := l.P1, l.P2

	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	dx := x2 - x1
	dy := y2 - y1

	steps := math.Max(math.Abs(dx), math.Abs(dy))

	xIncrement := dx / steps
	yIncrement := dy / steps

	x := x1
	y := y1

	for i := 0; float64(i) <= steps; i++ {
		cleanupPixel(pixels, IPoint{X: round(x), Y: round(y + 1)}, c)
		setPixel(pixels, IPoint{X: round(x), Y: round(y)}, c)
		cleanupPixel(pixels, IPoint{X: round(x), Y: round(y - 1)}, c)
		x += xIncrement
		y += yIncrement
	}
}
