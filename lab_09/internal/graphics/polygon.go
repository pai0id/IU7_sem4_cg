package graphics

import "image/color"

type Polygon struct {
	Points []FPoint
}

func DrawPolygon(pixels *SafePixels, pol Polygon, col color.Color) {
	for i := 0; i < len(pol.Points)-1; i++ {
		drawLine(pixels, Line{pol.Points[i], pol.Points[i+1]}, col)
	}
	drawLine(pixels, Line{pol.Points[len(pol.Points)-1], pol.Points[0]}, col)
}
