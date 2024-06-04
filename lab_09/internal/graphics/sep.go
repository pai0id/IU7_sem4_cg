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

func crossProduct(p1, p2, p3 FPoint) float64 {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

func IsConvex(p Polygon) int {
	n := len(p.Points)
	if n < 3 {
		return 0
	}

	initialSign := 0
	for i := 0; i < n; i++ {
		a := p.Points[i]
		b := p.Points[(i+1)%n]
		c := p.Points[(i+2)%n]
		cp := crossProduct(a, b, c)
		if cp != 0 {
			if initialSign == 0 {
				initialSign = sign(cp)
			} else if sign(cp) != initialSign {
				return 0
			}
		}
	}
	return initialSign
}

func vector(v1, v2 Line) float64 {
	x1 := v1.P1.X - v1.P2.X
	y1 := v1.P1.Y - v1.P2.Y
	x2 := v2.P1.X - v2.P2.X
	y2 := v2.P1.Y - v2.P2.Y
	return x1*y2 - x2*y1
}

func isVisiable(point, peak1, peak2 FPoint, norm int) bool {
	v := vector(Line{point, peak1}, Line{peak2, peak1})
	return float64(norm)*v < 0
}

func isIntersection(ed1, ed2 Line, norm int) *FPoint {
	vis1 := isVisiable(ed1.P1, ed2.P1, ed2.P2, norm)
	vis2 := isVisiable(ed1.P2, ed2.P1, ed2.P2, norm)

	if (vis1 && !vis2) || (!vis1 && vis2) {
		p1, p2 := ed1.P1, ed1.P2
		q1, q2 := ed2.P1, ed2.P2

		delta := (p2.X-p1.X)*(q1.Y-q2.Y) - (q1.X-q2.X)*(p2.Y-p1.Y)
		deltaT := (q1.X-p1.X)*(q1.Y-q2.Y) - (q1.X-q2.X)*(q1.Y-p1.Y)

		if math.Abs(delta) <= 1e-6 {
			return &p2
		}

		t := deltaT / delta

		return &FPoint{
			X: ed1.P1.X + (ed1.P2.X-ed1.P1.X)*t,
			Y: ed1.P1.Y + (ed1.P2.Y-ed1.P1.Y)*t,
		}
	}
	return nil
}

func SutherlandHodgman(pixels *SafePixels, sep, pol Polygon, norm int, col color.Color) {
	var s, f FPoint

	for i := 0; i < len(sep.Points)-1; i++ {
		newPol := Polygon{}
		for j, point := range pol.Points {
			if j == 0 {
				f = point
			} else {
				t := isIntersection(Line{s, point}, Line{sep.Points[i], sep.Points[i+1]}, norm)
				if t != nil {
					newPol.Points = append(newPol.Points, *t)
				}
			}

			s = point
			if isVisiable(s, sep.Points[i], sep.Points[i+1], norm) {
				newPol.Points = append(newPol.Points, s)
			}
		}

		if len(newPol.Points) != 0 {
			t := isIntersection(Line{s, f}, Line{sep.Points[i], sep.Points[i+1]}, norm)
			if t != nil {
				newPol.Points = append(newPol.Points, *t)
			}
		}

		pol = newPol
	}

	if len(pol.Points) == 0 {
		return
	}

	DrawPolygon(pixels, pol, col)
}
