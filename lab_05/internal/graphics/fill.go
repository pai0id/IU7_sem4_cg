package graphics

func reversePixel(dst []IPoint, x, y int) []IPoint {
	for i := range dst {
		if dst[i].X == x && dst[i].Y == y {
			dst[i] = dst[len(dst)-1]
			return dst[:len(dst)-1]
		}
	}
	return append(dst, IPoint{x, y})
}

func Fill(fig []FPoint) []IPoint {
	var points []IPoint

	return points
}

func FillWDelay(fig []FPoint) [][]IPoint {
	var points_states [][]IPoint

	return points_states
}
