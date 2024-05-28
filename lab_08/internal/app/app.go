package app

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"lab_08/internal/graphics"
)

var COLORS = []string{
	"Белый",
	"Черный",
	"Красный",
	"Зеленый",
	"Синий",
}

var (
	bgColorV        color.Color = getColor(COLORS[0])
	areaColorV      color.Color = getColor(COLORS[2])
	linePrimeColorV color.Color = getColor(COLORS[1])
	lineSepColorV   color.Color = getColor(COLORS[3])
)

var (
	raster   *canvas.Raster
	myWindow fyne.Window
	buttons  []*widget.Button
)

var pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}

var (
	currPoly graphics.Polygon = graphics.Polygon{}
	inPoly   bool             = false
)

var (
	currLines []graphics.Line
	inLine    bool = false
)

type tappableCanvasObject struct {
	fyne.CanvasObject
	OnTapped          func(x, y float64)
	OnTappedSecondary func(x, y float64)
}

func addRectDot(p graphics.FPoint) {
	if inPoly {
		prev := currPoly.Points[len(currPoly.Points)-1]
		if prev == p {
			dialog.ShowError(errors.New("ТА ЖЕ ТОЧКА"), myWindow)
			log.Println("Error: Та же точка")
			return
		}
		currPoly.Points = append(currPoly.Points, p)
		graphics.DrawLine(&pixels, graphics.Line{P1: p, P2: prev}, areaColorV)
		raster.Refresh()
	} else {
		clearCanvas()
		currPoly.Points = append(currPoly.Points, p)

		inPoly = true
		lockButtons()
	}
}

func linkPoly() {
	if len(currPoly.Points) < 3 {
		dialog.ShowError(errors.New("НЕТ ТРЕХ ТОЧЕК"), myWindow)
		log.Println("Error: Нет трех точек")
		return
	}
	graphics.DrawLine(&pixels, graphics.Line{P1: currPoly.Points[0], P2: currPoly.Points[len(currPoly.Points)-1]}, areaColorV)
	raster.Refresh()

	inPoly = false
	unlockButtons()
}

func addLineDot(p graphics.FPoint) {
	if inLine {
		if currLines[len(currLines)-1].P1 == p {
			dialog.ShowError(errors.New("ТА ЖЕ ТОЧКА"), myWindow)
			log.Println("Error: Та же точка")
			return
		}
		currLines[len(currLines)-1].P2 = p
		graphics.DrawLine(&pixels, currLines[len(currLines)-1], linePrimeColorV)
		raster.Refresh()

		inLine = false
		if !inPoly {
			unlockButtons()
		}
	} else {
		currLines = append(currLines, graphics.Line{P1: p})

		inLine = true
		lockButtons()
	}
}

func addParallelLines() {
	if len(currPoly.Points) < 3 || inPoly {
		dialog.ShowError(errors.New("НЕТ ПОЛИГОНА"), myWindow)
		log.Println("Error: Нет пролигона")
		return
	}

	for i := range currPoly.Points {
		var p1, p2, newP1, newP2 graphics.FPoint
		var newLine graphics.Line

		if i != len(currPoly.Points)-1 {
			p1 = currPoly.Points[i]
			p2 = currPoly.Points[i+1]
		} else {
			p1 = currPoly.Points[i]
			p2 = currPoly.Points[0]
		}

		newP1 = graphics.FPoint{X: p1.X - 10, Y: p1.Y - 10}
		newP2 = graphics.FPoint{X: p2.X - 10, Y: p2.Y - 10}
		newLine = graphics.Line{P1: newP1, P2: newP2}
		currLines = append(currLines, newLine)
		graphics.DrawLine(&pixels, newLine, linePrimeColorV)

		newP1 = graphics.FPoint{X: p1.X + 10, Y: p1.Y + 10}
		newP2 = graphics.FPoint{X: p2.X + 10, Y: p2.Y + 10}
		newLine = graphics.Line{P1: newP1, P2: newP2}
		currLines = append(currLines, newLine)
		graphics.DrawLine(&pixels, newLine, linePrimeColorV)
	}

	raster.Refresh()
}

func MakeTappable(canvas fyne.CanvasObject, onTapped func(x, y float64), onTappedSecondary func(x, y float64)) *tappableCanvasObject {
	return &tappableCanvasObject{CanvasObject: canvas, OnTapped: onTapped, OnTappedSecondary: onTappedSecondary}
}

func (t *tappableCanvasObject) Tapped(ev *fyne.PointEvent) {
	t.OnTapped(float64(ev.Position.X), float64(ev.Position.Y))
}

func (t *tappableCanvasObject) TappedSecondary(ev *fyne.PointEvent) {
	t.OnTappedSecondary(float64(ev.Position.X), float64(ev.Position.Y))
}

func (t *tappableCanvasObject) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.CanvasObject)
}

func CanvasOnTapped(x, y float64) {
	point := graphics.FPoint{X: x * float64(myWindow.Canvas().Scale()), Y: y * float64(myWindow.Canvas().Scale())}
	addLineDot(point)
}

func CanvasOnTappedSecondary(x, y float64) {
	point := graphics.FPoint{X: x * float64(myWindow.Canvas().Scale()), Y: y * float64(myWindow.Canvas().Scale())}
	addRectDot(point)
}

func lockButtons() {
	for _, v := range buttons {
		v.Disable()
	}
}

func unlockButtons() {
	for _, v := range buttons {
		v.Enable()
	}
}

func getColor(name string) color.Color {
	switch name {
	case COLORS[0]:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case COLORS[1]:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	case COLORS[2]:
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	case COLORS[3]:
		return color.RGBA{R: 0, G: 255, B: 0, A: 255}
	case COLORS[4]:
		return color.RGBA{R: 0, G: 0, B: 255, A: 255}
	default:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}
}

func drawCanvas(x, y, _, _ int) color.Color {
	var res color.Color
	pixels.MU.Lock()
	if v, keyExists := pixels.PXS[graphics.IPoint{X: x, Y: y}]; keyExists {
		res = v
	} else {
		res = bgColorV
	}
	pixels.MU.Unlock()
	return res
}

func clearCanvas() {
	pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}
	currLines = []graphics.Line{}
	currPoly = graphics.Polygon{}
	raster.Refresh()
}

func createVertSepRect(height float32) *canvas.Rectangle {
	sepRect := canvas.NewRectangle(theme.ForegroundColor())
	sepRect.SetMinSize(fyne.NewSize(1, height))
	return sepRect
}

func createHoriSepRect(width float32) *canvas.Rectangle {
	sepRect := canvas.NewRectangle(theme.ForegroundColor())
	sepRect.SetMinSize(fyne.NewSize(width, 1))
	return sepRect
}

func SetupApp() {
	myApp := app.New()
	myWindow = myApp.NewWindow("Geometry")

	methodLabel1 := canvas.NewText("Алгоритм", theme.ForegroundColor())
	methodLabel1.Alignment = fyne.TextAlignCenter
	methodLabel2 := canvas.NewText("Кируса-Бека", theme.ForegroundColor())
	methodLabel2.Alignment = fyne.TextAlignCenter

	methodLabelC := container.NewVBox(methodLabel1, methodLabel2)

	methodC := container.NewCenter(methodLabelC)

	sepRect0 := createVertSepRect(methodC.MinSize().Height)

	areaColorLabel := canvas.NewText("Выберите цвет прямоугольника", theme.ForegroundColor())
	areaColorSelect := widget.NewSelect(COLORS, func(value string) {
		areaColorV = getColor(value)
		log.Println("Select set to", value)
	})
	areaColorSelect.SetSelected(COLORS[2])
	areaColorV = getColor(COLORS[2])
	areaColorC := container.NewHBox(areaColorLabel, areaColorSelect)

	linePrimeColorLabel := canvas.NewText("Выберите цвет отрезков", theme.ForegroundColor())
	linePrimeColorSelect := widget.NewSelect(COLORS, func(value string) {
		linePrimeColorV = getColor(value)
		log.Println("Select set to", value)
	})
	linePrimeColorSelect.SetSelected(COLORS[1])
	linePrimeColorV = getColor(COLORS[1])
	linePrimeColorC := container.NewHBox(linePrimeColorLabel, linePrimeColorSelect)

	lineSepColorLabel := canvas.NewText("Выберите цвет отсечений", theme.ForegroundColor())
	lineSepColorSelect := widget.NewSelect(COLORS, func(value string) {
		lineSepColorV = getColor(value)
		log.Println("Select set to", value)
	})
	lineSepColorSelect.SetSelected(COLORS[3])
	lineSepColorV = getColor(COLORS[3])
	lineSepColorC := container.NewHBox(lineSepColorLabel, lineSepColorSelect)

	colorC := container.NewVBox(areaColorC, linePrimeColorC, lineSepColorC)

	sepRect1 := createVertSepRect(methodC.MinSize().Height)

	sepButton := widget.NewButton("Отсечь", func() {
		if len(currPoly.Points) < 3 || inPoly {
			dialog.ShowError(errors.New("НЕТ ПОЛИГОНА"), myWindow)
			log.Println("Error: Нет пролигона")
			return
		}
		norm := graphics.IsConvex(currPoly)
		if norm == 0 {
			dialog.ShowError(errors.New("ОТСЕКАТЕЛЬ НЕ ВЫПУКЛЫЙ"), myWindow)
			log.Println("Error: Не выпуклый")
			return
		}
		for _, l := range currLines {
			graphics.CyrusBeck(&pixels, currPoly, norm, l, lineSepColorV)
		}
		raster.Refresh()
	})
	buttons = append(buttons, sepButton)

	clearCanvasButton := widget.NewButton("Очистить экран", clearCanvas)
	buttons = append(buttons, clearCanvasButton)

	linkPolyButton := widget.NewButton("Замкнуть", linkPoly)

	addParallelLinesButton := widget.NewButton("Добавить параллельные линии", addParallelLines)
	buttons = append(buttons, addParallelLinesButton)

	canvasButtonsC := container.NewVBox(sepButton, clearCanvasButton, linkPolyButton, addParallelLinesButton)

	upperFrame := container.NewHBox(methodC, sepRect0, colorC, sepRect1, canvasButtonsC)

	raster = canvas.NewRasterWithPixels(drawCanvas)
	raster.SetMinSize(fyne.NewSize(upperFrame.MinSize().Width, myWindow.Canvas().Scale()*500))

	rasterC := container.NewPadded(MakeTappable(raster, CanvasOnTapped, CanvasOnTappedSecondary))
	rasterSizeLabel := canvas.NewText("", theme.ForegroundColor())

	lowerFrame := container.NewVBox(rasterC, rasterSizeLabel)

	sepRectULWFrames := createHoriSepRect(upperFrame.MinSize().Width)

	content := container.NewVBox(upperFrame, sepRectULWFrames, lowerFrame)

	go func() {
		for {
			time.Sleep(time.Second)
			w := int(raster.Size().Width * myWindow.Canvas().Scale())
			h := int(raster.Size().Height * myWindow.Canvas().Scale())
			rasterSizeLabel.Text = fmt.Sprintf("%dx%d", w, h)
			raster.Refresh()
		}
	}()
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
