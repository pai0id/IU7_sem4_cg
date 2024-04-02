package calc

import (
	"math"
)

func CircleCanon(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)

	lim := xC + IntNum(radius/math.Sqrt2) + 1
	for x := xC; x < lim; x++ {
		y := IntNum(math.Sqrt(math.Pow(radius, 2)-math.Pow(float64(x-xC), 2)) + yCenter)
		points = OctDup(points, x, y, xC, yC)
	}

	return points
}

func CircleParam(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)

	step := 1.0 / radius

	for t := 0.0; t <= math.Pi/4+step; t += step {
		x := xC + IntNum(radius*math.Cos(t))
		y := yC + IntNum(radius*math.Sin(t))
		points = OctDup(points, x, y, xC, yC)
	}

	return points
}

func CircleBres(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)
	r := IntNum(radius)

	x := 0
	y := IntNum(radius)
	delta := 2 * (1 - r)

	points = OctDup(points, x+xC, y+yC, xC, yC)

	for x < y {
		if delta <= 0 {
			deltaTemp := 2*(delta+y) - 1
			x++
			if deltaTemp >= 0 {
				delta += 2 * (x - y + 1)
				y--
			} else {
				delta += 2*x + 1
			}

		} else {
			deltaTemp := 2*(delta-x) - 1
			y--
			if deltaTemp < 0 {
				delta += 2 * (x - y + 1)
				x++
			} else {
				delta -= 2*y - 1
			}
		}

		points = OctDup(points, x+xC, y+yC, xC, yC)
	}

	return points
}

func CircleMidPoint(xCenter, yCenter, radius float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)
	r := IntNum(radius)

	x := r
	y := 0

	points = OctDup(points, x+xC, y+yC, xC, yC)
	delta := 1 - r

	for x > y {
		y++
		if delta > 0 {
			x--
			delta -= 2*x - 2
		}
		delta += 2*y + 3
		points = OctDup(points, x+xC, y+yC, xC, yC)
	}

	return points
}
