package graphics

import (
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
	PXS map[IPoint]bool
}

func round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}

func itofPoint(p IPoint) FPoint {
	return FPoint{float64(p.X), float64(p.Y)}
}

func ftoiPoint(p FPoint) IPoint {
	return IPoint{round(p.X), round(p.Y)}
}
