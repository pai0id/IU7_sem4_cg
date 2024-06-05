package graphics

import (
	"image/color"
	"math"
)

type Line struct {
	P1, P2 FPoint
}

type vec struct {
	X, Y float64
}

func GetVec(p1, p2 FPoint) vec {
	return vec{X: p2.X - p1.X, Y: p2.Y - p1.Y}
}

func GetCos(l1, l2 Line) float64 {
	v1 := GetVec(l1.P1, l1.P2)
	v2 := GetVec(l2.P1, l2.P2)

	up := v1.X*v2.X + v1.Y*v2.Y
	down := math.Sqrt(v1.X*v1.X+v1.Y*v1.Y) * math.Sqrt(v2.X*v2.X+v2.Y*v2.Y)

	return up / down
}

func GetLen(l Line) float64 {
	return math.Sqrt(math.Pow(l.P2.X-l.P1.X, 2) + math.Pow(l.P2.Y-l.P1.Y, 2))
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
