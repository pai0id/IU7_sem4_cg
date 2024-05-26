package graphics

import (
	"image/color"
	"math"
)

type (
	sepRectType [4]float64
	codeType    [4]int
)

func getCode(p FPoint, r sepRectType) codeType {
	var code codeType

	if p.X < r[0] {
		code[0] = 1
	}
	if p.X > r[1] {
		code[1] = 1
	}
	if p.Y < r[2] {
		code[2] = 1
	}
	if p.Y > r[3] {
		code[3] = 1
	}

	return code
}

func sumCode(code codeType) int {
	var res int

	for i := range code {
		res += code[i]
	}

	return res
}

func logProd(code1, code2 codeType) int {
	var p int = 0

	for i := range code1 {
		p += code1[i] & code2[i]
	}

	return p
}

func isVisible(l Line, r sepRectType) int {
	code1 := getCode(l.P1, r)
	code2 := getCode(l.P2, r)

	s1 := sumCode(code1)
	s2 := sumCode(code2)

	if s1 == 0 && s2 == 0 {
		return 1
	}
	log := logProd(code1, code2)

	if log != 0 {
		return 0
	}

	return 2
}

func getSepRect(r Rect) sepRectType {
	var res sepRectType

	res[0] = math.Min(r.P1.X, r.P2.X)
	res[1] = math.Max(r.P1.X, r.P2.X)
	res[2] = math.Min(r.P1.Y, r.P2.Y)
	res[3] = math.Max(r.P1.Y, r.P2.Y)

	return res
}

func SutherlandCohen(pixels *SafePixels, r Rect, l Line, c color.Color) {
	o := getSepRect(r)
	var flag int = 0
	var angle float64

	if l.P1.X == l.P2.X {
		flag = -1
	} else {
		angle = (l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X)
		if angle == 0 {
			flag = 1
		}
	}

	for i := range o {
		vis := isVisible(l, o)

		if vis == 1 {
			break
		} else if vis == 0 {
			return
		}

		code1 := getCode(l.P1, o)
		code2 := getCode(l.P2, o)

		if code1[i] == code2[i] {
			continue
		}

		if code1[i] == 0 {
			l.P1, l.P2 = l.P2, l.P1
		}

		if flag != -1 {
			if i < 2 {
				l.P1.Y += angle * (o[i] - l.P1.X)
				l.P1.X = o[i]
				continue
			} else {
				l.P1.X += (1 / angle) * (o[i] - l.P1.Y)
			}
		}

		l.P1.Y = o[i]
	}

	LineCDA(pixels, l, c)
}
