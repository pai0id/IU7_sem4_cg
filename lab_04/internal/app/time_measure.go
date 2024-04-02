package app

import (
	"image/color"
	"log"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var NTESTS int = 500

type Plot struct {
	xVals []float64
	yVals []float64
}

func drawPlots(plots []Plot, xName string) string {
	dir := "data"
	os.Mkdir(dir, os.ModePerm)

	p := plot.New()
	p.X.Label.Text = xName
	p.Y.Label.Text = "Время (нс)"

	colors := []color.Color{
		color.RGBA{R: 0x66, G: 0xc2, B: 0xa5, A: 0xff},
		color.RGBA{R: 0xfc, G: 0x8d, B: 0x62, A: 0xff},
		color.RGBA{R: 0x8d, G: 0xa0, B: 0xcb, A: 0xff},
		color.RGBA{R: 0xe7, G: 0x8a, B: 0xc3, A: 0xff},
		color.RGBA{R: 0xa6, G: 0xd8, B: 0x54, A: 0xff},
	}

	for i, plot := range plots {
		points := []plotter.XY{}
		for i, v := range plot.xVals {
			points = append(points, plotter.XY{X: v, Y: plot.yVals[i]})
		}
		lp, err := plotter.NewLine(plotter.XYs(points))
		if err != nil {
			log.Fatalf("Could not create line plotter: %v", err)
			return ""
		}

		lp.LineStyle.Color = colors[i%len(colors)]

		p.Add(lp)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "data/plots.png"); err != nil {
		log.Fatalf("Could not save plot: %v", err)
		return ""
	}

	return "data/plots.png"
}

func MeasureCircle(minRad, maxRad, step float64) string {
	var plots []Plot

	for _, fName := range METHODS {
		f := parseCircMethod(fName)
		var plot Plot

		for r := minRad; r <= maxRad; r += step {
			var t int64 = 0
			for i := 0; i < NTESTS; i++ {
				st := time.Now()
				f(0, 0, r)
				t += time.Since(st).Nanoseconds()
			}

			var el float64 = float64(t) / float64(NTESTS)

			plot.xVals = append(plot.xVals, r)
			plot.yVals = append(plot.yVals, el)
		}
		plots = append(plots, plot)
	}
	return drawPlots(plots, "Радиус")
}

func MeasureEllipse(minHeight, minWidth, step float64, cnt int64) string {
	var plots []Plot

	for _, fName := range METHODS {
		f := parseEllipseMethod(fName)
		var plot Plot

		w := minWidth
		h := minHeight
		for k := int64(0); k < cnt; k++ {
			var t int64 = 0
			for i := 0; i < NTESTS; i++ {
				st := time.Now()
				f(0, 0, w, h)
				t += time.Since(st).Nanoseconds()
			}

			var el float64 = float64(t) / float64(NTESTS)

			plot.xVals = append(plot.xVals, w)
			plot.yVals = append(plot.yVals, el)

			w += step
			h += step
		}
		plots = append(plots, plot)
	}
	return drawPlots(plots, "Ширина")
}
