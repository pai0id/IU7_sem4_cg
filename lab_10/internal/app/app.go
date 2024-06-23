package app

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"lab_10/internal/graphics"
)

var (
	bgColorV   color.Color = color.White
	drawColorV color.Color = color.Black
)

var (
	raster   *canvas.Raster
	myWindow fyne.Window
)

func drawCanvas(x, y, _, _ int) color.Color {
	var res color.Color
	graphics.Pixels.MU.Lock()
	if _, keyExists := graphics.Pixels.PXS[graphics.IPoint{X: x, Y: y}]; keyExists {
		res = drawColorV
	} else {
		res = bgColorV
	}
	graphics.Pixels.MU.Unlock()
	return res
}

func clearCanvas() {
	graphics.Pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]bool)}
	go raster.Refresh()
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

	funcLabel := canvas.NewText("	Функция		", theme.ForegroundColor())
	funcLabel.Alignment = fyne.TextAlignCenter
	funcSelect := widget.NewSelect(graphics.FuncStrArr, func(value string) {
		graphics.CurrFunc = graphics.GetID(value)
		log.Println("Select set to", value)
		clearCanvas()
		graphics.Solve()
		go raster.Refresh()
	})
	funcSelect.SetSelected(graphics.FuncStrArr[0])
	graphics.CurrFunc = graphics.GetID(graphics.FuncStrArr[0])
	funcC := container.NewVBox(funcLabel, funcSelect)

	sepRect0 := createVertSepRect(funcC.MinSize().Height)

	emptyLabel := canvas.NewText("	", theme.ForegroundColor())
	emptyLabel.Alignment = fyne.TextAlignCenter

	fromLabel := canvas.NewText("От", theme.ForegroundColor())
	fromLabel.Alignment = fyne.TextAlignCenter

	toLabel := canvas.NewText("До", theme.ForegroundColor())
	toLabel.Alignment = fyne.TextAlignCenter

	stepLabel := canvas.NewText("Шаг", theme.ForegroundColor())
	stepLabel.Alignment = fyne.TextAlignCenter

	xLabel := canvas.NewText("x", theme.ForegroundColor())
	xLabel.Alignment = fyne.TextAlignCenter

	zLabel := canvas.NewText("z", theme.ForegroundColor())
	zLabel.Alignment = fyne.TextAlignCenter

	xFromEntry := widget.NewEntry()
	xFromEntry.MultiLine = false
	xFromEntry.SetText("-10")

	zFromEntry := widget.NewEntry()
	zFromEntry.MultiLine = false
	zFromEntry.SetText("-10")

	xToEntry := widget.NewEntry()
	xToEntry.MultiLine = false
	xToEntry.SetText("10")

	zToEntry := widget.NewEntry()
	zToEntry.MultiLine = false
	zToEntry.SetText("10")

	xStepEntry := widget.NewEntry()
	xStepEntry.MultiLine = false
	xStepEntry.SetText("0.1")

	zStepEntry := widget.NewEntry()
	zStepEntry.MultiLine = false
	zStepEntry.SetText("0.1")

	fromToStepButton := widget.NewButton("Применить", func() {
		xFrom, err := strconv.ParseFloat(xFromEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		xTo, err := strconv.ParseFloat(xToEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		xStep, err := strconv.ParseFloat(xStepEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
			log.Println("Error:", err)
			return
		}
		zFrom, err := strconv.ParseFloat(zFromEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Z"), myWindow)
			log.Println("Error:", err)
			return
		}
		zTo, err := strconv.ParseFloat(zToEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Z"), myWindow)
			log.Println("Error:", err)
			return
		}
		zStep, err := strconv.ParseFloat(zStepEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
			log.Println("Error:", err)
			return
		}
		clearCanvas()
		graphics.SetMeta(xFrom, xTo, xStep, zFrom, zTo, zStep)
		go raster.Refresh()
	})

	fromToStepGrid := container.NewGridWithColumns(4,
		emptyLabel, fromLabel, toLabel, stepLabel,
		xLabel, xFromEntry, xToEntry, xStepEntry,
		zLabel, zFromEntry, zToEntry, zStepEntry)

	fromToStepC := container.NewVBox(fromToStepGrid, fromToStepButton)

	sepRect1 := createVertSepRect(funcC.MinSize().Height)

	coefLabel := canvas.NewText("Коэф. масштабирования", theme.ForegroundColor())
	coefLabel.Alignment = fyne.TextAlignCenter

	coefEntry := widget.NewEntry()
	coefEntry.MultiLine = false
	coefEntry.SetText("50")

	coefButton := widget.NewButton("Изменить", func() {
		coef, err := strconv.Atoi(coefEntry.Text)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ КОЭФИЦИЕНТ"), myWindow)
			log.Println("Error:", err)
			return
		}
		clearCanvas()
		graphics.SetSF(coef)
		go raster.Refresh()
	})

	xRotateLabel := canvas.NewText("x", theme.ForegroundColor())
	xRotateLabel.Alignment = fyne.TextAlignCenter

	xRotateEntry := widget.NewEntry()
	xRotateEntry.MultiLine = false
	xRotateEntry.SetText("45")

	xRotateButton := widget.NewButton("Вращать", func() {
		xRotate, err := strconv.ParseFloat(xRotateEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ УГОЛ"), myWindow)
			log.Println("Error:", err)
			return
		}
		clearCanvas()
		graphics.RotateX(xRotate)
		go raster.Refresh()
	})

	yRotateLabel := canvas.NewText("y", theme.ForegroundColor())
	yRotateLabel.Alignment = fyne.TextAlignCenter

	yRotateEntry := widget.NewEntry()
	yRotateEntry.MultiLine = false
	yRotateEntry.SetText("45")

	yRotateButton := widget.NewButton("Вращать", func() {
		yRotate, err := strconv.ParseFloat(yRotateEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ УГОЛ"), myWindow)
			log.Println("Error:", err)
			return
		}
		clearCanvas()
		graphics.RotateY(yRotate)
		go raster.Refresh()
	})

	zRotateLabel := canvas.NewText("z", theme.ForegroundColor())
	zRotateLabel.Alignment = fyne.TextAlignCenter

	zRotateEntry := widget.NewEntry()
	zRotateEntry.MultiLine = false
	zRotateEntry.SetText("45")

	zRotateButton := widget.NewButton("Вращать", func() {
		zRotate, err := strconv.ParseFloat(zRotateEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ УГОЛ"), myWindow)
			log.Println("Error:", err)
			return
		}
		clearCanvas()
		graphics.RotateZ(zRotate)
		go raster.Refresh()
	})

	scaleRotateGrid := container.NewGridWithColumns(3,
		coefLabel, coefEntry, coefButton,
		xRotateLabel, xRotateEntry, xRotateButton,
		yRotateLabel, yRotateEntry, yRotateButton,
		zRotateLabel, zRotateEntry, zRotateButton)

	sepRect2 := createVertSepRect(funcC.MinSize().Height)

	drawButton := widget.NewButton("Нарисовать", func() {
		clearCanvas()
		graphics.Solve()
		go raster.Refresh()
	})

	clearCanvasButton := widget.NewButton("Очистить экран", clearCanvas)

	buttonsC := container.NewVBox(drawButton, clearCanvasButton)

	upperFrame := container.NewHBox(funcC, sepRect0, fromToStepC, sepRect1, scaleRotateGrid, sepRect2, buttonsC)

	raster = canvas.NewRasterWithPixels(drawCanvas)
	raster.SetMinSize(fyne.NewSize(upperFrame.MinSize().Width, myWindow.Canvas().Scale()*500))

	rasterSizeLabel := canvas.NewText("", theme.ForegroundColor())

	lowerFrame := container.NewVBox(raster, rasterSizeLabel)

	sepRectULWFrames := createHoriSepRect(upperFrame.MinSize().Width)

	content := container.NewVBox(upperFrame, sepRectULWFrames, lowerFrame)

	go func() {
		time.Sleep(time.Second)
		w := int(raster.Size().Width * myWindow.Canvas().Scale())
		h := int(raster.Size().Height * myWindow.Canvas().Scale())
		rasterSizeLabel.Text = fmt.Sprintf("%dx%d", w, h)
		if graphics.UpdateScreen(w, h) {
			clearCanvas()
			graphics.Solve()
			go raster.Refresh()
		}
	}()
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
