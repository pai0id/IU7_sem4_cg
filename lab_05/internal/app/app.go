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

	"lab_05/internal/graphics"
)

type Dot struct {
	point   graphics.FPoint
	fig_num int
	end     *Dot
	filled  bool
}

var COLORS = []string{
	"Белый",
	"Черный",
	"Красный",
	"Зеленый",
	"Синий",
}

var bgColorV color.Color
var colorV color.Color

var raster *canvas.Raster
var myWindow fyne.Window

var pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}

var dots = []Dot{}
var dotsTable *widget.Table
var currFig int = 1

type tappableCanvasObject struct {
	fyne.CanvasObject
	OnTapped          func(x, y float64)
	OnTappedSecondary func()
}

func addDot(p graphics.FPoint) {
	dot := Dot{point: p, fig_num: currFig, end: nil, filled: false}
	if len(dots) != 0 && currFig == dots[len(dots)-1].fig_num {
		if dot.point == dots[len(dots)-1].point {
			dialog.ShowError(errors.New("ТА ЖЕ ТОЧКА"), myWindow)
			log.Println("Error: Та же точка")
			return
		}
		getPixels(graphics.LineCDA(dot.point, dots[len(dots)-1].point))
		raster.Refresh()
	}
	dots = append(dots, dot)
	dotsTable.ScrollToBottom()
	dotsTable.Refresh()
}

func closeFig() {
	cnt := 0
	for _, v := range dots {
		if v.fig_num == currFig {
			cnt++
		}
		if cnt > 2 || v.fig_num > currFig {
			break
		}
	}
	if cnt <= 2 {
		dialog.ShowError(errors.New("НЕЛЬЗЯ ЗАКРЫТЬ ФИГУРУ"), myWindow)
		log.Println("Error: Нельзя закрыть фигуру")
		return
	}

	for _, v := range dots {
		if v.fig_num == currFig {
			getPixels(graphics.LineCDA(v.point, dots[len(dots)-1].point))
			dots[len(dots)-1].end = &v
			raster.Refresh()
			break
		}
	}
	currFig++
}

func MakeTappable(canvas fyne.CanvasObject, onTapped func(x, y float64), onTappedSecondary func()) *tappableCanvasObject {
	return &tappableCanvasObject{CanvasObject: canvas, OnTapped: onTapped, OnTappedSecondary: onTappedSecondary}
}

func (t *tappableCanvasObject) Tapped(ev *fyne.PointEvent) {
	t.OnTapped(float64(ev.Position.X), float64(ev.Position.Y))
}

func (t *tappableCanvasObject) TappedSecondary(ev *fyne.PointEvent) {
	t.OnTappedSecondary()
}

func (t *tappableCanvasObject) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.CanvasObject)
}

func CanvasOnTapped(x, y float64) {
	addDot(graphics.FPoint{X: x * float64(myWindow.Canvas().Scale()), Y: y * float64(myWindow.Canvas().Scale())})
}

func CanvasOnTappedSecondary() {
	closeFig()
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
	if colr, keyExists := pixels.PXS[graphics.IPoint{X: x, Y: y}]; keyExists {
		res = colr
	} else {
		res = bgColorV
	}
	pixels.MU.Unlock()
	return res
}

func clearCanvas() {
	pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}
	dots = nil
	currFig = 1
	raster.Refresh()
}

func drawFigs() {
	if len(dots) == 0 {
		return
	}
	for i := range dots {
		if i == 0 {
			continue
		}
		if dots[i-1].end != nil {
			continue
		}
		if dots[i].end != nil {
			getPixels(graphics.LineCDA(dots[i].point, dots[i-1].point))
			getPixels(graphics.LineCDA(dots[i].point, dots[i].end.point))
			continue
		}
		getPixels(graphics.LineCDA(dots[i].point, dots[i-1].point))
	}
}

func clearFills() {
	pixels = graphics.SafePixels{PXS: make(map[graphics.IPoint]color.Color)}
	for i := range dots {
		dots[i].filled = false
	}
	drawFigs()
	raster.Refresh()
}

func getPixels(points []graphics.IPoint) {
	pixels.MU.Lock()
	for _, point := range points {
		pixels.PXS[point] = colorV
	}
	pixels.MU.Unlock()
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
	// myApp.Settings().SetTheme(theme.LightTheme())

	methodLabel1 := canvas.NewText("Алгоритм", theme.ForegroundColor())
	methodLabel1.Alignment = fyne.TextAlignCenter
	methodLabel2 := canvas.NewText("заполнения", theme.ForegroundColor())
	methodLabel2.Alignment = fyne.TextAlignCenter
	methodLabel3 := canvas.NewText("по ребрам", theme.ForegroundColor())
	methodLabel3.Alignment = fyne.TextAlignCenter

	methodLabelC := container.NewVBox(methodLabel1, methodLabel2, methodLabel3)

	methodC := container.NewCenter(methodLabelC)

	sepRect0 := createVertSepRect(methodC.MinSize().Height)

	bgColorLabel := canvas.NewText("Выберите цвет фона", theme.ForegroundColor())

	bgColorSelect := widget.NewSelect(COLORS, func(value string) {
		bgColorV = getColor(value)
		raster.Refresh()
		log.Println("Select set to", value)
	})
	bgColorSelect.SetSelected(COLORS[1])
	bgColorV = getColor(COLORS[1])

	colorLabel := canvas.NewText("Выберите цвет фигуры", theme.ForegroundColor())

	colorSelect := widget.NewSelect(COLORS, func(value string) {
		colorV = getColor(value)
		log.Println("Select set to", value)
	})
	colorSelect.SetSelected(COLORS[0])
	colorV = getColor(COLORS[0])

	colorC := container.NewVBox(bgColorLabel, bgColorSelect, colorLabel, colorSelect)

	sepRect1 := createVertSepRect(methodC.MinSize().Height)

	dotXLabel := canvas.NewText("Введите x\t", theme.ForegroundColor())

	dotXEntry := widget.NewEntry()
	dotXEntry.MultiLine = false
	dotXEntry.SetText("200")

	dotYLabel := canvas.NewText("Введите y\t", theme.ForegroundColor())

	dotYEntry := widget.NewEntry()
	dotYEntry.MultiLine = false
	dotYEntry.SetText("200")

	dotC := container.NewVBox(dotXLabel, dotXEntry, dotYLabel, dotYEntry)

	sepRect2 := createVertSepRect(methodC.MinSize().Height)

	dotButton := widget.NewButton("Добавить точку", func() {
		x, err := strconv.ParseFloat(dotXEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		y, err := strconv.ParseFloat(dotYEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Y"), myWindow)
			log.Println("Error:", err)
			return
		}
		addDot(graphics.FPoint{X: x, Y: y})
	})

	closeButton := widget.NewButton("Замкнуть", func() {
		closeFig()
	})

	clearFillsButton := widget.NewButton("Очистить заполнения", clearFills)

	clearCanvasButton := widget.NewButton("Очистить экран", clearCanvas)

	canvasButtonsC := container.NewVBox(dotButton, closeButton, clearFillsButton, clearCanvasButton)

	sepRect3 := createVertSepRect(methodC.MinSize().Height)

	dotsTable = widget.NewTable(
		func() (int, int) {
			return len(dots) + 1, 3
		},
		func() fyne.CanvasObject {
			o := widget.NewLabel("		")
			o.Alignment = fyne.TextAlignCenter
			return o
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				if i.Col == 0 {
					o.(*widget.Label).SetText("№ Фигуры")
				} else if i.Col == 1 {
					o.(*widget.Label).SetText("X")
				} else {
					o.(*widget.Label).SetText("Y")
				}
			} else {
				if i.Col == 0 {
					o.(*widget.Label).SetText(fmt.Sprintf("%d", dots[i.Row-1].fig_num))
				} else if i.Col == 1 {
					o.(*widget.Label).SetText(fmt.Sprintf("%.2f", dots[i.Row-1].point.X))
				} else {
					o.(*widget.Label).SetText(fmt.Sprintf("%.2f", dots[i.Row-1].point.Y))
				}
			}
		})

	tableC := container.NewGridWrap(fyne.NewSize(300, canvasButtonsC.MinSize().Height), dotsTable)

	sepRect4 := createVertSepRect(methodC.MinSize().Height)

	fillButton := widget.NewButton("Заполнить", func() {
		var fig []graphics.FPoint
		for i, v := range dots {
			fig = append(fig, v.point)
			if v.end != nil {
				if v.filled {
					fig = nil
					continue
				}
				graphics.Fill(&pixels, fig, colorV)
				raster.Refresh()
				fig = nil
				dots[i].filled = true
			}
		}
		tmp := colorV
		colorV = getColor(COLORS[2])
		drawFigs()
		colorV = tmp
	})

	fillSleepButton := widget.NewButton("Заполнить с задержкой", func() {
		var fig []graphics.FPoint
		for i, v := range dots {
			fig = append(fig, v.point)
			if v.end != nil {
				if v.filled {
					fig = nil
					continue
				}
				c := make(chan int)
				cnt := 0
				go graphics.FillWDelay(&pixels, fig, colorV, c)
				for n := range c {
					cnt += n
					raster.Refresh()
					time.Sleep(time.Millisecond * 10)
				}
				fig = nil
				dots[i].filled = true
			}
		}
		tmp := colorV
		colorV = getColor(COLORS[2])
		drawFigs()
		colorV = tmp
	})

	fillTimeButton := widget.NewButton("Замерить время", func() {
		var fig []graphics.FPoint
		var t_all int64 = 0
		for i, v := range dots {
			fig = append(fig, v.point)
			if v.end != nil {
				if v.filled {
					fig = nil
					continue
				}
				st := time.Now()
				graphics.Fill(&pixels, fig, colorV)
				t_all += time.Since(st).Nanoseconds()
				raster.Refresh()
				fig = nil
				dots[i].filled = true
			}
		}
		tmp := colorV
		colorV = getColor(COLORS[2])
		drawFigs()
		colorV = tmp
		dialog.ShowInformation("Результат", fmt.Sprintf("Затрачено %d нс", t_all), myWindow)
	})

	fillButtonsC := container.NewVBox(fillButton, fillSleepButton, fillTimeButton)

	upperFrame := container.NewHBox(methodC, sepRect0, colorC, sepRect1, dotC, sepRect2, canvasButtonsC, sepRect3, tableC, sepRect4, fillButtonsC)

	raster = canvas.NewRasterWithPixels(drawCanvas)
	raster.SetMinSize(fyne.NewSize(upperFrame.MinSize().Width, myWindow.Canvas().Scale()*500))

	rasterSizeLabel := canvas.NewText(fmt.Sprintf("%dx%d", int(raster.MinSize().Width), int(raster.MinSize().Height)), theme.ForegroundColor())

	rasterC := container.NewPadded(MakeTappable(raster, CanvasOnTapped, CanvasOnTappedSecondary))

	lowerFrame := container.NewVBox(rasterC, rasterSizeLabel)

	sepRectULWFrames := createHoriSepRect(upperFrame.MinSize().Width)

	content := container.NewVBox(upperFrame, sepRectULWFrames, lowerFrame)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
