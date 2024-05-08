package graphics

import (
	"image/color"
	"sync"
)

type IPoint struct {
	X, Y int
}

type FPoint struct {
	X, Y float64
}

type SafePixels struct {
	MU  sync.Mutex
	PXS map[IPoint]color.Color
}

func round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}
