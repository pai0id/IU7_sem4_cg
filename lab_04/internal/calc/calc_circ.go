package calc

import (
	"math"
)

func CircleCanon(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))

	lim := int(math.Round(xCenter + radius/math.Sqrt2))
	for x := xC; x <= lim; x++ {
		y := int(math.Round(math.Sqrt(radius*radius-math.Pow(float64(x)-xCenter, 2)) + yCenter))
		points = OctDup(points, x, y, xC, yC)
	}

	return points
}

func CircleParam(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))

	step := 1.0 / radius

	for t := 0.0; t <= math.Pi/4+step; t += step {
		x := xC + int(math.Round(radius*math.Cos(t)))
		y := yC + int(math.Round(radius*math.Sin(t)))
		points = OctDup(points, x, y, xC, yC)
	}

	return points
}

func CircleBres(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))
	r := int(math.Round(radius))

	x := 0
	y := r
	d := 3 - (2 * r)

	points = OctDup(points, x+xC, y+yC, xC, yC)

	for x <= y {
		x++

		if d < 0 {
			d = d + (4 * x) + 6
		} else {
			y--
			d = d + 4*(x-y) + 10
		}

		points = OctDup(points, x+xC, y+yC, xC, yC)
	}

	return points
}

func CircleMidPoint(xCenter, yCenter, radius float64) []Point {
	var points []Point

	// xC := int(math.Round(xCenter)
	// yC := int(math.Round(yCenter)
	// r := int(math.Round(radius)

	return points
}
