package calc

import (
	"math"
)

func EllipseCanon(xCenter, yCenter, width, height float64) []Point {
	var points []Point
	var lim int

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)

	lim = IntNum(xCenter + width/math.Sqrt(1+math.Pow(height, 2)/math.Pow(width, 2)))
	for x := xC; x <= lim; x++ {
		y := IntNum(math.Sqrt(math.Pow(width, 2)*math.Pow(height, 2)-math.Pow(float64(x)-xCenter, 2)*math.Pow(height, 2))/width + yCenter)
		points = QuaDup(points, x, y, xC, yC)
	}

	lim = IntNum(yCenter + height/math.Sqrt(1+math.Pow(width, 2)/math.Pow(height, 2)))
	for y := lim; y > yC; y-- {
		x := IntNum(math.Sqrt(math.Pow(width, 2)*math.Pow(height, 2)-math.Pow(float64(y)-yCenter, 2)*math.Pow(width, 2))/height + xCenter)
		points = QuaDup(points, x, y, xC, yC)
	}

	return points
}

func EllipseParam(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)

	step := 1.0 / math.Max(width, height)

	for t := 0.0; t <= math.Pi/2+step; t += step {
		x := xC + IntNum(width*math.Cos(t))
		y := yC + IntNum(height*math.Sin(t))
		points = QuaDup(points, x, y, xC, yC)
	}

	return points
}

func EllipseBres(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)
	w := IntNum(width)
	h := IntNum(height)

	x := 0
	y := h
	delta := h*h - w*w*(2*h+1)

	points = QuaDup(points, x+xC, y+yC, xC, yC)

	for y > 0 {
		if delta <= 0 {
			deltaTemp := 2*delta + w*w*(2*y-1)
			x++
			delta += h * h * (2*x + 1)
			if deltaTemp >= 0 {
				y--
				delta += w * w * (-2*y + 1)
			}

		} else {
			deltaTemp := 2*delta + h*h*(-2*x-1)
			y--
			delta += w * w * (-2*y + 1)
			if deltaTemp < 0 {
				x++
				delta += h * h * (2*x + 1)
			}
		}

		points = QuaDup(points, x+xC, y+yC, xC, yC)
	}

	return points
}

func EllipseMidPoint(xCenter, yCenter, width, height float64) []Point {
	var points []Point

	xC := IntNum(xCenter)
	yC := IntNum(yCenter)

	x := 0.0
	y := height

	delta := height*height - width*width*height + 0.25*width*width
	dx := 2 * height * height * x
	dy := 2 * width * width * y

	for dx > dy {
		points = QuaDup(points, IntNum(x+xCenter), IntNum(y+yCenter), xC, yC)

		x++
		dx += 2 * height * height

		if delta >= 0 {
			y--
			dy -= 2 * width * width
			delta -= dy
		}

		delta += dx + height*height
	}

	delta = height*height*(x+0.5)*(x+0.5) + width*width*(y-1)*(y-1) - width*width*height*height

	for y >= 0 {
		points = QuaDup(points, IntNum(x+xCenter), IntNum(y+yCenter), xC, yC)

		y--
		dy -= 2 * width * width

		if delta <= 0 {
			x++
			dx += 2 * height * height
			delta += dx
		}

		delta -= dy - width*width
	}

	return points
}
