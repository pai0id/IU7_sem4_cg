package graphics

type IPoint struct {
	X, Y int
}

type FPoint struct {
	X, Y float64
}

func round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}
