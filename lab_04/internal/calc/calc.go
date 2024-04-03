package calc

type Point struct {
	X, Y int
}

func OctDup(points []Point, x, y, xCenter, yCenter int) []Point {
	points = QuaDup(points, x, y, xCenter, yCenter)
	points = QuaDup(points, y+xCenter-yCenter, x+yCenter-xCenter, xCenter, yCenter)

	return points
}

func QuaDup(points []Point, x, y, xCenter, yCenter int) []Point {
	points = append(points, Point{x, y})
	points = append(points, Point{2*xCenter - x, y})
	points = append(points, Point{x, 2*yCenter - y})
	points = append(points, Point{2*xCenter - x, 2*yCenter - y})

	return points
}
