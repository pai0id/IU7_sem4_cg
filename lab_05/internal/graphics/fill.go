package graphics

import (
	"image/color"
)

func revPixel(pixels *SafePixels, point IPoint, col color.Color) {
	(*pixels).MU.Lock()
	if _, keyExists := (*pixels).PXS[point]; keyExists {
		delete((*pixels).PXS, point)
	} else {
		(*pixels).PXS[point] = col
	}
	(*pixels).MU.Unlock()
}

func Fill(pixels *SafePixels, fig []FPoint, col color.Color) {
	xm := round(fig[0].X)
	for _, p := range fig {
		if p.X > float64(xm) {
			xm = round(p.X)
		}
	}

	for i := range fig {
		var p1, p2 FPoint
		if i == 0 {
			p1 = fig[len(fig)-1]
			p2 = fig[0]
		} else {
			p1 = fig[i-1]
			p2 = fig[i]
		}

		var ed [2]FPoint
		if p1.Y == p2.Y {
			continue
		} else if p1.Y > p2.Y {
			ed = [2]FPoint{p2, p1}
		} else {
			ed = [2]FPoint{p1, p2}
		}

		y := ed[0].Y
		endY := ed[1].Y

		dx := ed[1].X - ed[0].X
		dy := ed[1].Y - ed[0].Y

		xIncrement := dx / dy

		startX := ed[0].X

		for y < endY-1 {
			x := round(startX)
			for x < xm {
				revPixel(pixels, IPoint{x, round(y)}, col)
				x++
			}

			startX += xIncrement
			y++
		}
	}
}

func FillWDelay(pixels *SafePixels, fig []FPoint, col color.Color, c chan int) {
	var n int = 0

	xm := round(fig[0].X)
	for _, p := range fig {
		if p.X > float64(xm) {
			xm = round(p.X)
		}
	}

	for i := range fig {
		var p1, p2 FPoint
		if i == 0 {
			p1 = fig[len(fig)-1]
			p2 = fig[0]
		} else {
			p1 = fig[i-1]
			p2 = fig[i]
		}

		var ed [2]FPoint
		if p1.Y == p2.Y {
			continue
		} else if p1.Y > p2.Y {
			ed = [2]FPoint{p2, p1}
		} else {
			ed = [2]FPoint{p1, p2}
		}

		y := ed[0].Y
		endY := ed[1].Y

		dx := ed[1].X - ed[0].X
		dy := ed[1].Y - ed[0].Y

		xIncrement := dx / dy

		startX := ed[0].X

		for y < endY-1 {
			x := round(startX)
			for x < xm {
				revPixel(pixels, IPoint{x, round(y)}, col)
				x++
			}
			if n%2 == 0 {
				c <- n
			}
			n++
			startX += xIncrement
			y++
		}
	}
	close(c)
}
