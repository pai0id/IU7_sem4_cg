package app

import (
	"image/color"
	"lab_05/internal/graphics"
	"time"
)

var NTESTS int64 = 5000

func FillMeasureTime(pixels *graphics.SafePixels, fig []graphics.FPoint, col color.Color) float64 {
	var testArr = []int64{}
	var avr float64
	var t int64 = 0
	var i int64 = 0
	for ; i < NTESTS; i++ {
		st := time.Now()
		graphics.Fill(pixels, fig, col)
		t += time.Since(st).Nanoseconds()
		testArr = append(testArr, time.Since(st).Nanoseconds())

		if i%100 == 99 {
			avr = float64(t)/float64(i) + 1.0
			if calcRSE(i+1, avr, calcStdev(testArr, avr)) <= 1 {
				i++
				break
			}
		}
	}

	avr = float64(t) / float64(i)
	return avr
}
