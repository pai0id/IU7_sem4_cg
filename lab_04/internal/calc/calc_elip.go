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

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))
	w := int(math.Round(width))
	h := int(math.Round(height))
	sqw := w * w
	sqh := h * h

	x := 0
	y := h
	delta := 4*sqh*(x+1)*(x+1) + sqw*((2*y-1)*(2*y-1)) - 4*sqw*sqh

	for sqw*(2*y-1) > 2*sqh*(x+1) {
		points = QuaDup(points, x+xC, y+yC, xC, yC)
		if delta < 0 {
			x++
			delta += 4 * h * h * (2*x + 3)
		} else {
			x++
			delta = delta - 8*sqw*(y-1) + 4*sqh*(2*x+3)
			y--
		}
	}

	delta = sqh*((2*x+1)*(2*x+1)) + 4*sqw*((y+1)*(y+1)) - 4*sqw*sqh

	for y+1 != 0 {
		points = QuaDup(points, x+xC, y+yC, xC, yC)
		if delta < 0 {
			y--
			delta += 4 * w * w * (2*y + 3)
		} else {
			y--
			delta = delta - 8*sqh*(x+1) + 4*sqw*(2*y+3)
			x++
		}
	}

	return points
}

func EllipseMidPoint(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	xC := int(math.Round(xCenter))
	yC := int(math.Round(yCenter))
	w := int(math.Round(width))
	h := int(math.Round(height))
	sqw := w * w
	sqh := h * h

	x := 0
	y := h

	points = QuaDup(points, x+xC, y+yC, xC, yC)

	border := int(math.Round(float64(w) / math.Sqrt(1+float64(sqh)/float64(sqw))))
	delta := sqh - int(math.Round(float64(sqw)*(float64(h)-0.25)))

	for x <= border {
		if delta < 0 {
			x++
			delta += 2*sqh*x + 1
		} else {
			x++
			y--
			delta += 2*sqh*x - 2*sqw*y + 1
		}

		points = QuaDup(points, x+xC, y+yC, xC, yC)
	}

	x = w
	y = 0

	points = QuaDup(points, x+xC, y+yC, xC, yC)

	border = int(math.Round(float64(h) / math.Sqrt(1+float64(sqw)/float64(sqh))))
	delta = sqw - int(math.Round(float64(sqh)*(float64(w)-0.25)))

	for y <= border {
		if delta < 0 {
			y++
			delta += 2*sqw*y + 1
		} else {
			x--
			y++
			delta += 2*sqw*y - 2*sqh*x + 1
		}

		points = QuaDup(points, x+xC, y+yC, xC, yC)
	}

	return points
}
