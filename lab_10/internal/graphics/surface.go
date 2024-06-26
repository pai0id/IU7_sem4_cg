package graphics

import (
	"math"
)

var (
	sf          = 50
	xFrom       = -10.0
	xTo         = 10.0
	xStep       = 0.1
	zFrom       = -10.0
	zTo         = 10.0
	zStep       = 0.1
	transMatrix = identityMatrix()
)

var (
	screenWidth  int
	screenHeight int
)

func UpdateScreen(w, h int) bool {
	if screenWidth != w || screenHeight != h {
		screenWidth = w
		screenHeight = h
		return true
	}
	return false
}

func identityMatrix() [][]float64 {
	matrix := make([][]float64, 4)
	for i := range matrix {
		matrix[i] = make([]float64, 4)
		matrix[i][i] = 1.0
	}
	return matrix
}

func rotateTransMatrix(rotateMatrix [][]float64) {
	resMatrix := make([][]float64, 4)
	for i := range resMatrix {
		resMatrix[i] = make([]float64, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				resMatrix[i][j] += transMatrix[i][k] * rotateMatrix[k][j]
			}
		}
	}
	transMatrix = resMatrix
}

func transPoint(point []float64) []float64 {
	point = append(point, 1)
	resPoint := make([]float64, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			resPoint[i] += point[j] * transMatrix[j][i]
		}
	}
	for i := 0; i < 3; i++ {
		resPoint[i] *= float64(sf)
	}
	resPoint[0] += float64(screenWidth) / 2
	resPoint[1] += float64(screenHeight) / 2
	return resPoint[:3]
}

func RotateX(angle float64) {
	value := angle / 180 * math.Pi
	rotateMatrix := [][]float64{
		{1, 0, 0, 0},
		{0, math.Cos(value), math.Sin(value), 0},
		{0, -math.Sin(value), math.Cos(value), 0},
		{0, 0, 0, 1},
	}
	rotateTransMatrix(rotateMatrix)
	Solve()
}

func RotateY(angle float64) {
	value := angle / 180 * math.Pi
	rotateMatrix := [][]float64{
		{math.Cos(value), 0, -math.Sin(value), 0},
		{0, 1, 0, 0},
		{math.Sin(value), 0, math.Cos(value), 0},
		{0, 0, 0, 1},
	}
	rotateTransMatrix(rotateMatrix)
	Solve()
}

func RotateZ(angle float64) {
	value := angle / 180 * math.Pi
	rotateMatrix := [][]float64{
		{math.Cos(value), math.Sin(value), 0, 0},
		{-math.Sin(value), math.Cos(value), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	rotateTransMatrix(rotateMatrix)
	Solve()
}

func SetSF(scale int) {
	sf = scale
	Solve()
}

func SetMeta(xF, xT, xS, zF, zT, zS float64) {
	xFrom = xF
	xTo = xT
	xStep = xS
	zFrom = zF
	zTo = zT
	zStep = zS
	Solve()
}

func isVisible(point FPoint) bool {
	return point.X >= 0 && point.X < float64(screenWidth) && point.Y >= 0 && point.Y < float64(screenHeight)
}

func drawPoint(x, y int, hh, lh []int) bool {
	p := FPoint{X: float64(x), Y: float64(y)}
	if !isVisible(p) {
		return false
	}
	if y > hh[x] {
		hh[x] = y
		setPixel(IPoint{x, y})
	} else if y < lh[x] {
		lh[x] = y
		setPixel(IPoint{x, y})
	}
	return true
}

func drawHorizonPart(p1, p2 FPoint, hh, lh []int) {
	if p1.X > p2.X {
		p1, p2 = p2, p1
	}
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	l := math.Max(math.Abs(dx), math.Abs(dy))
	dx /= l
	dy /= l
	x, y := p1.X, p1.Y
	for i := 0; i <= int(l); i++ {
		if !drawPoint(int(math.Round(x)), int(math.Round(y)), hh, lh) {
			return
		}
		x += dx
		y += dy
	}
}

func drawHorizon(f func(float64, float64) float64, hh, lh []int, fr, to, step, z float64) {
	var prev *FPoint
	for x := fr; x <= to; x += step {
		current := FPointFromSlice(transPoint([]float64{x, f(x, z), z}))
		if prev != nil {
			drawHorizonPart(*prev, current, hh, lh)
		}
		prev = &current
	}
}

func Solve() {
	f := CurrFunc
	highHorizon := make([]int, screenWidth)
	lowHorizon := make([]int, screenWidth)
	for i := range lowHorizon {
		lowHorizon[i] = screenHeight
	}
	for z := zFrom; z <= zTo; z += zStep {
		drawHorizon(f, highHorizon, lowHorizon, xFrom, xTo, xStep, z)
	}
	for z := zFrom; z < zTo; z += zStep {
		p1 := FPointFromSlice(transPoint([]float64{xFrom, f(xFrom, z), z}))
		p2 := FPointFromSlice(transPoint([]float64{xFrom, f(xFrom, z+zStep), z + zStep}))
		l := Line{P1: p1, P2: p2}
		go DrawLine(l)
		p1 = FPointFromSlice(transPoint([]float64{xTo, f(xTo, z), z}))
		p2 = FPointFromSlice(transPoint([]float64{xTo, f(xTo, z+zStep), z + zStep}))
		l = Line{P1: p1, P2: p2}
		go DrawLine(l)
	}
}
