package graphics

import "math"

func LineCDA(p1, p2 FPoint) []IPoint {
	var points []IPoint

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
		points = append(points, IPoint{X: round(x), Y: round(y)})
		x += xIncrement
		y += yIncrement
	}

	return points
}
