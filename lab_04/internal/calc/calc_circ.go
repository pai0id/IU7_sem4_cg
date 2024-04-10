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
	d := 1 - 2*r
	e := 0

	for x <= y {
		points = OctDup(points, x+xC, y+yC, xC, yC)

		e = 2*(d+y) - 1

		if (d < 0) && (e <= 0) {
			x++
			d += 2*x + 1
		} else if (d > 0) && (e > 0) {
			y--
			d -= 2*y + 1
		} else {
			x++
			y--
			d += 2 * (x - y)
		}
	}

	return points
}

func CircleMidPoint(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))
	r := int(math.Round(radius))

	x := r
	y := 0
	p := 1 - r

	for x >= y {
		points = OctDup(points, x+xC, y+yC, xC, yC)

		y++
		if p < 0 {
			p += 2*y + 1
		} else {
			x--
			p += 2*y - 2*x + 1
		}
	}

	return points
}
