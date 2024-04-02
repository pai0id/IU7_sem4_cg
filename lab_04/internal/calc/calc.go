package calc

type Point struct {
	X, Y int
}

func IntNum(num float64) int {
	if num > 0 {
		return int(num + 0.5)
	} else {
		return int(num - 0.5)
	}
}

func OctDup(points []Point, x, y, xCenter, yCenter int) []Point {
	points = append(points, Point{x, y})
	points = append(points, Point{2*xCenter - x, y})
	points = append(points, Point{x, 2*yCenter - y})
	points = append(points, Point{2*xCenter - x, 2*yCenter - y})
	points = append(points, Point{y + xCenter - yCenter, x + yCenter - xCenter})
	points = append(points, Point{-y + xCenter + yCenter, x + yCenter - xCenter})
	points = append(points, Point{y + xCenter - yCenter, -x + yCenter + xCenter})
	points = append(points, Point{-y + xCenter + yCenter, -x + yCenter + xCenter})

	return points
}

func QuaDup(points []Point, x, y, xCenter, yCenter int) []Point {
	points = append(points, Point{x, y})
	points = append(points, Point{2*xCenter - x, y})
	points = append(points, Point{x, 2*yCenter - y})
	points = append(points, Point{2*xCenter - x, 2*yCenter - y})

	return points
}
