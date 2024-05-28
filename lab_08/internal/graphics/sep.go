package graphics

import (
	"image/color"
	"math"
)

func sign(x float64) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func IsConvex(p Polygon) int {
	flag := 1

	vo := p.Points[0]
	vi := p.Points[1]
	vn := p.Points[2]

	x1 := vi.X - vo.X
	y1 := vi.Y - vo.Y

	x2 := vn.X - vi.X
	y2 := vn.Y - vi.Y

	r := x1*y2 - x2*y1
	prev := sign(r)

	for i := 2; i < len(p.Points)-1; i++ {
		if flag == 0 {
			break
		}
		vo = p.Points[i-1]
		vi = p.Points[i]
		vn = p.Points[i+1]

		x1 = vi.X - vo.X
		y1 = vi.Y - vo.Y

		x2 = vn.X - vi.X
		y2 = vn.Y - vi.Y

		r = x1*y2 - x2*y1
		curr := sign(r)

		if curr != prev {
			flag = 0
		}
		prev = curr
	}

	vo = p.Points[len(p.Points)-1]
	vi = p.Points[0]
	vn = p.Points[1]

	x1 = vi.X - vo.X
	y1 = vi.Y - vo.Y

	x2 = vn.X - vi.X
	y2 = vn.Y - vi.Y

	r = x1*y2 - x2*y1
	curr := sign(r)
	if curr != prev {
		flag = 0
	}

	return flag * curr
}

func scalar(v1, v2 FPoint) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func CyrusBeck(pixels *SafePixels, p Polygon, n int, l Line, c color.Color) {
	tb := 0.0
	te := 1.0

	D := FPoint{X: l.P2.X - l.P1.X, Y: l.P2.Y - l.P1.Y}

	for i := 0; i < len(p.Points); i++ {
		W := FPoint{X: l.P1.X - p.Points[i].X, Y: l.P1.Y - p.Points[i].Y}

		N := FPoint{}
		if i == len(p.Points)-1 {
			N.X = -float64(n) * (p.Points[0].Y - p.Points[i].Y)
			N.Y = float64(n) * (p.Points[0].X - p.Points[i].X)
		} else {
			N.X = -float64(n) * (p.Points[i+1].Y - p.Points[i].Y)
			N.Y = float64(n) * (p.Points[i+1].X - p.Points[i].X)
		}

		Dscalar := scalar(D, N)
		Wscalar := scalar(W, N)

		if Dscalar == 0 {
			if Wscalar < 0 {
				return
			}
		} else {
			t := -Wscalar / Dscalar
			if Dscalar > 0 {
				if t > 1 {
					return
				}
				tb = math.Max(tb, t)
			} else {
				if t < 0 {
					return
				}
				te = math.Min(te, t)
			}
		}
	}

	if tb <= te {
		drawLine(pixels, Line{
			P1: FPoint{X: l.P1.X + (l.P2.X-l.P1.X)*te, Y: l.P1.Y + (l.P2.Y-l.P1.Y)*te},
			P2: FPoint{X: l.P1.X + (l.P2.X-l.P1.X)*tb, Y: l.P1.Y + (l.P2.Y-l.P1.Y)*tb}}, c)
	}
}
