package calc

import (
	"math"
)

func EllipseCanon(xCenter, yCenter, width, height float64) []Point {
	var points []Point
	var lim int

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))

	lim = int(math.Round(xCenter + width/math.Sqrt2))
	for x := xC; x <= lim+xC; x++ {
		y := int(math.Round(math.Sqrt(math.Pow(height, 2)*(1.0-math.Pow(float64(x)-xCenter, 2)/math.Pow(width, 2))) + yCenter))
		points = QuaDup(points, x, y, xC, yC)
	}
	lim = int(math.Round(xCenter + height/math.Sqrt2))
	for y := yC; y <= lim+yC; y++ {
		x := int(math.Round(math.Sqrt(math.Pow(width, 2)*(1.0-math.Pow(float64(y)-yCenter, 2)/math.Pow(height, 2))) + xCenter))
		points = QuaDup(points, x, y, xC, yC)
	}

	return points
}

func EllipseParam(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))

	step := 1.0 / math.Max(width, height)

	for t := 0.0; t <= math.Pi/2+step; t += step {
		x := xC + int(math.Round(width*math.Cos(t)))
		y := yC + int(math.Round(height*math.Sin(t)))
		points = QuaDup(points, x, y, xC, yC)
	}

	return points
}

func EllipseBres(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	// xC := int(math.Round(xCenter)
	// yC := int(math.Round(yCenter)
	// w := int(math.Round(width)
	// h := int(math.Round(height)

	return points
}

func EllipseMidPoint(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	// xC := int(math.Round(xCenter)
	// yC := int(math.Round(yCenter)

	return points
}
