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

	"lab_09/internal/graphics"
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
	sepColorV       color.Color = getColor(COLORS[2])
	polyPrimeColorV color.Color = getColor(COLORS[1])
	polySepColorV   color.Color = getColor(COLORS[3])
)

var (
	raster     *canvas.Raster
	myWindow   fyne.Window
	buttons    []*widget.Button
	linkButton *widget.Button
)

var pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}

var (
	currSep graphics.Polygon = graphics.Polygon{}
	inSep   bool             = false
)

var (
	currPoly graphics.Polygon = graphics.Polygon{}
	inPoly   bool             = false
)

type tappableCanvasObject struct {
	fyne.CanvasObject
	OnTapped          func(x, y float64)
	OnTappedSecondary func(x, y float64)
}

func addSepDot(p graphics.FPoint) {
	if inSep {
		prev := currSep.Points[len(currSep.Points)-1]
		if prev == p {
			dialog.ShowError(errors.New("ТА ЖЕ ТОЧКА"), myWindow)
			log.Println("Error: Та же точка")
			return
		}
		currSep.Points = append(currSep.Points, p)
		graphics.DrawLine(&pixels, graphics.Line{P1: p, P2: prev}, sepColorV)
		raster.Refresh()
	} else {
		clearSep()
		currSep.Points = append(currSep.Points, p)

		inSep = true
		lockButtons()
		linkButton.Enable()
	}
}

func linkSep() {
	if len(currSep.Points) < 3 {
		dialog.ShowError(errors.New("НЕТ ТРЕХ ТОЧЕК"), myWindow)
		log.Println("Error: Нет трех точек")
		return
	}
	graphics.DrawLine(&pixels, graphics.Line{P1: currSep.Points[0], P2: currSep.Points[len(currSep.Points)-1]}, sepColorV)
	raster.Refresh()

	currSep.Points = append(currSep.Points, currSep.Points[0])
	inSep = false
	unlockButtons()
	linkButton.Disable()
}

func addPolyDot(p graphics.FPoint) {
	if inPoly {
		prev := currPoly.Points[len(currPoly.Points)-1]
		if prev == p {
			dialog.ShowError(errors.New("ТА ЖЕ ТОЧКА"), myWindow)
			log.Println("Error: Та же точка")
			return
		}
		currPoly.Points = append(currPoly.Points, p)
		graphics.DrawLine(&pixels, graphics.Line{P1: p, P2: prev}, polyPrimeColorV)
		raster.Refresh()
	} else {
		clearPoly()
		currPoly.Points = append(currPoly.Points, p)

		inPoly = true
		lockButtons()
		linkButton.Enable()
	}
}

func linkPoly() {
	if len(currPoly.Points) < 3 {
		dialog.ShowError(errors.New("НЕТ ТРЕХ ТОЧЕК"), myWindow)
		log.Println("Error: Нет трех точек")
		return
	}
	graphics.DrawLine(&pixels, graphics.Line{P1: currPoly.Points[0], P2: currPoly.Points[len(currPoly.Points)-1]}, polyPrimeColorV)
	raster.Refresh()

	currPoly.Points = append(currPoly.Points, currPoly.Points[0])
	inPoly = false
	unlockButtons()
	linkButton.Disable()
}

func link() {
	if inPoly {
		linkPoly()
	} else if inSep {
		linkSep()
	}
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
	addPolyDot(point)
}

func CanvasOnTappedSecondary(x, y float64) {
	point := graphics.FPoint{X: x * float64(myWindow.Canvas().Scale()), Y: y * float64(myWindow.Canvas().Scale())}
	addSepDot(point)
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
	currPoly = graphics.Polygon{}
	currSep = graphics.Polygon{}
	raster.Refresh()
}

func clearPoly() {
	pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}
	currPoly = graphics.Polygon{}
	if len(currSep.Points) != 0 {
		graphics.DrawPolygon(&pixels, currSep, sepColorV)
	}
	raster.Refresh()
}

func clearSep() {
	pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}
	currSep = graphics.Polygon{}
	if len(currPoly.Points) != 0 {
		graphics.DrawPolygon(&pixels, currPoly, polyPrimeColorV)
	}
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
	methodLabel2 := canvas.NewText("САЗЕРЛЕНДА-ХОДЖМЕНА", theme.ForegroundColor())
	methodLabel2.Alignment = fyne.TextAlignCenter

	methodLabelC := container.NewVBox(methodLabel1, methodLabel2)

	methodC := container.NewCenter(methodLabelC)

	sepRect0 := createVertSepRect(methodC.MinSize().Height)

	areaColorLabel := canvas.NewText("Выберите цвет отсекателя", theme.ForegroundColor())
	areaColorSelect := widget.NewSelect(COLORS, func(value string) {
		sepColorV = getColor(value)
		log.Println("Select set to", value)
	})
	areaColorSelect.SetSelected(COLORS[2])
	sepColorV = getColor(COLORS[2])
	areaColorC := container.NewHBox(areaColorLabel, areaColorSelect)

	linePrimeColorLabel := canvas.NewText("Выберите цвет полигона", theme.ForegroundColor())
	linePrimeColorSelect := widget.NewSelect(COLORS, func(value string) {
		polyPrimeColorV = getColor(value)
		log.Println("Select set to", value)
	})
	linePrimeColorSelect.SetSelected(COLORS[1])
	polyPrimeColorV = getColor(COLORS[1])
	linePrimeColorC := container.NewHBox(linePrimeColorLabel, linePrimeColorSelect)

	lineSepColorLabel := canvas.NewText("Выберите цвет отсечения", theme.ForegroundColor())
	lineSepColorSelect := widget.NewSelect(COLORS, func(value string) {
		polySepColorV = getColor(value)
		log.Println("Select set to", value)
	})
	lineSepColorSelect.SetSelected(COLORS[3])
	polySepColorV = getColor(COLORS[3])
	lineSepColorC := container.NewHBox(lineSepColorLabel, lineSepColorSelect)

	colorC := container.NewVBox(areaColorC, linePrimeColorC, lineSepColorC)

	sepRect1 := createVertSepRect(methodC.MinSize().Height)

	sepButton := widget.NewButton("Отсечь", func() {
		if len(currSep.Points) < 3 || inSep {
			dialog.ShowError(errors.New("НЕТ ОТСЕКАТЕЛЯ"), myWindow)
			log.Println("Error: Нет отсекателя")
			return
		}
		if len(currPoly.Points) < 3 || inPoly {
			dialog.ShowError(errors.New("НЕТ ПОЛИГОНА"), myWindow)
			log.Println("Error: Нет полигона")
			return
		}
		norm := graphics.IsConvex(currSep)
		if norm == 0 {
			dialog.ShowError(errors.New("ОТСЕКАТЕЛЬ НЕ ВЫПУКЛЫЙ"), myWindow)
			log.Println("Error: Не выпуклый")
			return
		}
		graphics.SutherlandHodgman(&pixels, currSep, currPoly, norm, polySepColorV)
		raster.Refresh()
	})
	buttons = append(buttons, sepButton)

	clearCanvasButton := widget.NewButton("Очистить экран", clearCanvas)
	buttons = append(buttons, clearCanvasButton)

	linkButton = widget.NewButton("Замкнуть", link)
	linkButton.Disable()

	canvasButtonsC := container.NewVBox(sepButton, clearCanvasButton, linkButton)

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
