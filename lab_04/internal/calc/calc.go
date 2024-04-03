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

func QuaDup(points []Point, x, y, xCenter, yCenter int) []Point {
	points = append(points, Point{x, y})
	points = append(points, Point{2*xCenter - x, y})
	points = append(points, Point{x, 2*yCenter - y})
	points = append(points, Point{2*xCenter - x, 2*yCenter - y})

	return points
}
