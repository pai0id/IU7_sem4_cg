package app

import (
	"errors"
	"image/color"
	"log"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"lab_04/internal/calc"
)

type Pixel struct {
	point  calc.Point
	pColor color.Color
}

var METHODS = []string{
	"Каноническое уравнение",
	"Параметрическое уравнение",
	"Алгоритм Брезенхема",
	"Алгоритм средней точки",
}

var COLORS = []string{
	"Белый",
	"Черный",
	"Красный",
	"Зеленый",
	"Синий",
}

var STEP_OPTIONS = []string{
	"Высота",
	"Ширина",
	"Обе",
}

var TIME_MEASURE_OPTIONS = []string{
	"Для окружностей",
	"Для эллипсов",
}

var bgColorV color.Color
var colorV color.Color
var raster *canvas.Raster
var pixels = []Pixel{}

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
	for _, pixel := range pixels {
		if x == pixel.point.X && y == pixel.point.Y {
			return pixel.pColor
		}
	}
	return bgColorV
}

func clearCanvas() {
	pixels = nil
	raster.Refresh()
}

func removePixel(pixels []Pixel, i int) []Pixel {
	pixels[i] = pixels[len(pixels)-1]
	return pixels[:len(pixels)-1]
}

func getPixels(points []calc.Point) []Pixel {
	for _, point := range points {
		for i, pixel := range pixels {
			if pixel.point.X == point.X && pixel.point.Y == point.Y {
				pixels = removePixel(pixels, i)
				break
			}
		}
		pixels = append(pixels, Pixel{point, colorV})
	}
	return pixels
}

// func getPixels(points []calc.Point) []Pixel {
// 	for _, point := range points {
// 		for i := 0; i < len(pixels); i++ {
// 			if (pixels[i].point.X == point.X || pixels[i].point.X == point.X+1 || pixels[i].point.X == point.X-1) &&
// 				(pixels[i].point.Y == point.Y || pixels[i].point.Y == point.Y+1 || pixels[i].point.Y == point.Y-1) {
// 				pixels = removePixel(pixels, i)
// 				i--
// 			}
// 		}
// 	}
// 	for _, point := range points {
// 		pixels = append(pixels, Pixel{point, colorV})
// 	}
// 	return pixels
// }

func parseCircMethod(method string) func(xC float64, yC float64, r float64) []calc.Point {
	switch method {
	case METHODS[0]:
		return calc.CircleCanon
	case METHODS[1]:
		return calc.CircleParam
	case METHODS[2]:
		return calc.CircleBres
	case METHODS[3]:
		return calc.CircleMidPoint
	default:
		return calc.CircleCanon
	}
}

func parseEllipseMethod(method string) func(xC float64, yC float64, w float64, h float64) []calc.Point {
	switch method {
	case METHODS[0]:
		return calc.EllipseCanon
	case METHODS[1]:
		return calc.EllipseParam
	case METHODS[2]:
		return calc.EllipseBres
	case METHODS[3]:
		return calc.EllipseMidPoint
	default:
		return calc.EllipseCanon
	}
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
	myWindow := myApp.NewWindow("Geometry")
	// myApp.Settings().SetTheme(theme.LightTheme())

	methodLabel := canvas.NewText("Выберите метод", theme.ForegroundColor())
	methodLabel.TextSize = 12

	methodSelect := widget.NewSelect(METHODS, func(value string) {
		log.Println("Select set to", value)
	})
	methodSelect.SetSelected(METHODS[0])

	methodC := container.NewGridWrap(fyne.NewSize(200, 30), methodLabel, methodSelect)

	bgColorLabel := canvas.NewText("Выберите цвет фона", theme.ForegroundColor())
	bgColorLabel.TextSize = 12

	bgColorSelect := widget.NewSelect(COLORS, func(value string) {
		bgColorV = getColor(value)
		raster.Refresh()
		log.Println("Select set to", value)
	})
	bgColorSelect.SetSelected(COLORS[1])
	bgColorV = getColor(COLORS[1])

	colorLabel := canvas.NewText("Выберите цвет фигуры", theme.ForegroundColor())
	colorLabel.TextSize = 12

	colorSelect := widget.NewSelect(COLORS, func(value string) {
		colorV = getColor(value)
		log.Println("Select set to", value)
	})
	colorSelect.SetSelected(COLORS[0])
	colorV = getColor(COLORS[0])

	colorC := container.NewVBox(bgColorLabel, bgColorSelect, colorLabel, colorSelect)

	sepRect1 := createVertSepRect(methodC.MinSize().Height)

	centerXLabel := canvas.NewText("Введите x\t", theme.ForegroundColor())
	centerXLabel.TextSize = 12

	centerXEntry := widget.NewEntry()
	centerXEntry.MultiLine = false
	centerXEntry.SetText("600")

	centerYLabel := canvas.NewText("Введите y\t", theme.ForegroundColor())
	centerYLabel.TextSize = 12

	centerYEntry := widget.NewEntry()
	centerYEntry.MultiLine = false
	centerYEntry.SetText("400")

	centerC := container.NewVBox(centerXLabel, centerXEntry, centerYLabel, centerYEntry)

	sepRect2 := createVertSepRect(methodC.MinSize().Height)

	radLabel := canvas.NewText("Введите радиус", theme.ForegroundColor())
	radLabel.TextSize = 12

	radEntry := widget.NewEntry()
	radEntry.MultiLine = false
	radEntry.SetText("100")

	circleButton := widget.NewButton("Окружность", func() {
		xCenter, err := strconv.ParseFloat(centerXEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		yCenter, err := strconv.ParseFloat(centerYEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Y"), myWindow)
			log.Println("Error:", err)
			return
		}
		radius, err := strconv.ParseFloat(radEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ РАДИУС"), myWindow)
			log.Println("Error:", err)
			return
		}
		if radius <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ РАДИУС"), myWindow)
			log.Println("Error: radius <= 0")
			return
		}
		f := parseCircMethod(methodSelect.Selected)
		points := f(xCenter, yCenter, radius)
		pixels = getPixels(points)
		raster.Refresh()
	})

	circleC := container.NewVBox(radLabel, radEntry, circleButton)

	sepRect3 := createVertSepRect(methodC.MinSize().Height)

	heightLabel := canvas.NewText("Введите высоту", theme.ForegroundColor())
	heightLabel.TextSize = 12

	heightEntry := widget.NewEntry()
	heightEntry.MultiLine = false
	heightEntry.SetText("100")

	widthLabel := canvas.NewText("Введите ширину", theme.ForegroundColor())
	widthLabel.TextSize = 12

	widthEntry := widget.NewEntry()
	widthEntry.MultiLine = false
	widthEntry.SetText("150")

	ellipseButton := widget.NewButton("Эллипс", func() {
		xCenter, err := strconv.ParseFloat(centerXEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		yCenter, err := strconv.ParseFloat(centerYEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Y"), myWindow)
			log.Println("Error:", err)
			return
		}
		height, err := strconv.ParseFloat(heightEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ ВЫСОТА"), myWindow)
			log.Println("Error:", err)
			return
		}
		if height <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ ВЫСОТА"), myWindow)
			log.Println("Error: height <= 0")
			return
		}
		width, err := strconv.ParseFloat(widthEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ ШИРИНА"), myWindow)
			log.Println("Error:", err)
			return
		}
		if width <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ ШИРИНА"), myWindow)
			log.Println("Error: width <= 0")
			return
		}
		f := parseEllipseMethod(methodSelect.Selected)
		points := f(xCenter, yCenter, width, height)
		pixels = getPixels(points)
		raster.Refresh()
	})

	ellipseCL := container.NewVBox(heightLabel, heightEntry, widthLabel, widthEntry)
	ellipseC := container.NewHBox(ellipseCL, ellipseButton)

	circleMinRadLabel := canvas.NewText(" Начальный радиус ", theme.ForegroundColor())
	circleMinRadLabel.TextSize = 12

	circleMinRadEntry := widget.NewEntry()
	circleMinRadEntry.MultiLine = false
	circleMinRadEntry.SetText("5")

	circleMaxRadLabel := canvas.NewText(" Конечный радиус ", theme.ForegroundColor())
	circleMaxRadLabel.TextSize = 12

	circleMaxRadEntry := widget.NewEntry()
	circleMaxRadEntry.MultiLine = false
	circleMaxRadEntry.SetText("100")

	circleRangeCL := container.NewVBox(circleMinRadLabel, circleMinRadEntry, circleMaxRadLabel, circleMaxRadEntry)

	circleStepLabel := canvas.NewText(" Шаг изменения радиуса ", theme.ForegroundColor())
	circleStepLabel.TextSize = 12

	circleStepEntry := widget.NewEntry()
	circleStepEntry.MultiLine = false
	circleStepEntry.SetText("5")

	circleRangeButton := widget.NewButton("Спектр окружностей", func() {
		xCenter, err := strconv.ParseFloat(centerXEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		yCenter, err := strconv.ParseFloat(centerYEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Y"), myWindow)
			log.Println("Error:", err)
			return
		}
		minRad, err := strconv.ParseFloat(circleMinRadEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MIN R"), myWindow)
			log.Println("Error:", err)
			return
		}
		if minRad <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MIN R"), myWindow)
			log.Println("Error: minRad <= 0")
			return
		}
		maxRad, err := strconv.ParseFloat(circleMaxRadEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MAX R"), myWindow)
			log.Println("Error:", err)
			return
		}
		if maxRad <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MAX R"), myWindow)
			log.Println("Error: maxRad <= 0")
			return
		}
		step, err := strconv.ParseFloat(circleStepEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
			log.Println("Error:", err)
			return
		}
		if (step > 0 && minRad > maxRad) || (step < 0 && minRad < maxRad) || step == 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
			log.Println("Error: inf loop")
			return
		}
		f := parseCircMethod(methodSelect.Selected)
		var points []calc.Point
		if step > 0 {
			for r := minRad; r <= maxRad; r += step {
				points = f(xCenter, yCenter, r)
				pixels = getPixels(points)
			}
		} else {
			for r := minRad; r >= maxRad; r += step {
				points = f(xCenter, yCenter, r)
				pixels = getPixels(points)
			}
		}
		raster.Refresh()
	})

	circleRangeCR := container.NewVBox(circleStepLabel, circleStepEntry, circleRangeButton)

	circleRangeC := container.NewHBox(circleRangeCL, circleRangeCR)

	sepRect4 := createHoriSepRect(circleC.MinSize().Width)

	ellipseMinHeightLabel := canvas.NewText(" Начальная высота ", theme.ForegroundColor())
	ellipseMinHeightLabel.TextSize = 12

	ellipseMinHeightEntry := widget.NewEntry()
	ellipseMinHeightEntry.MultiLine = false
	ellipseMinHeightEntry.SetText("100")

	ellipseMinWidthLabel := canvas.NewText(" Начальная ширина ", theme.ForegroundColor())
	ellipseMinWidthLabel.TextSize = 12

	ellipseMinWidthEntry := widget.NewEntry()
	ellipseMinWidthEntry.MultiLine = false
	ellipseMinWidthEntry.SetText("10")

	ellipseRangeCL := container.NewVBox(ellipseMinHeightLabel, ellipseMinHeightEntry, ellipseMinWidthLabel, ellipseMinWidthEntry)

	stepOptionLabel := canvas.NewText("Изменяемая полуось", theme.ForegroundColor())
	stepOptionLabel.TextSize = 12

	stepOptionSelect := widget.NewSelect(STEP_OPTIONS, func(value string) {
		log.Println("Select set to", value)
	})
	stepOptionSelect.SetSelected(STEP_OPTIONS[2])

	ellipseStepLabel := canvas.NewText(" Шаг изменения ", theme.ForegroundColor())
	ellipseStepLabel.TextSize = 12

	ellipseStepEntry := widget.NewEntry()
	ellipseStepEntry.MultiLine = false
	ellipseStepEntry.SetText("10")

	ellipseCntLabel := canvas.NewText(" Количество эллипсов ", theme.ForegroundColor())
	ellipseCntLabel.TextSize = 12
	ellipseCntLabel.Alignment = fyne.TextAlignCenter

	ellipseCntEntry := widget.NewEntry()
	ellipseCntEntry.MultiLine = false
	ellipseCntEntry.SetText("10")

	ellipseRangeButton := widget.NewButton("Спектр эллипсов", func() {
		xCenter, err := strconv.ParseFloat(centerXEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ X"), myWindow)
			log.Println("Error:", err)
			return
		}
		yCenter, err := strconv.ParseFloat(centerYEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ Y"), myWindow)
			log.Println("Error:", err)
			return
		}
		minHeight, err := strconv.ParseFloat(ellipseMinHeightEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ВЫСОТА"), myWindow)
			log.Println("Error:", err)
			return
		}
		if minHeight <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ВЫСОТА"), myWindow)
			log.Println("Error: minHeight <= 0")
			return
		}
		minWidth, err := strconv.ParseFloat(ellipseMinWidthEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ШИРИНА"), myWindow)
			log.Println("Error:", err)
			return
		}
		if minWidth <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ШИРИНА"), myWindow)
			log.Println("Error: minWidth <= 0")
			return
		}
		step, err := strconv.ParseFloat(ellipseStepEntry.Text, 64)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
			log.Println("Error:", err)
			return
		}
		cnt, err := strconv.ParseInt(ellipseCntEntry.Text, 10, 32)
		if err != nil {
			dialog.ShowError(errors.New("НЕКОРРЕКТНОЕ КОЛИЧЕСТВО"), myWindow)
			log.Println("Error:", err)
			return
		}
		if cnt <= 0 {
			dialog.ShowError(errors.New("НЕКОРРЕКТНОЕ КОЛИЧЕСТВО"), myWindow)
			log.Println("Error: cnt <= 0")
			return
		}
		f := parseEllipseMethod(methodSelect.Selected)
		if stepOptionSelect.Selected == STEP_OPTIONS[0] {
			h := minHeight
			for i := 0; i < int(cnt); i++ {
				points := f(xCenter, yCenter, minWidth, h)
				pixels = getPixels(points)
				h += step
				if h <= 0 {
					dialog.ShowError(errors.New("ДОСТИГНУТ РАЗМЕР 0"), myWindow)
					log.Println("Error: size 0")
					break
				}
			}
		} else if stepOptionSelect.Selected == STEP_OPTIONS[1] {
			w := minWidth
			for i := 0; i < int(cnt); i++ {
				points := f(xCenter, yCenter, w, minHeight)
				pixels = getPixels(points)
				w += step
				if w <= 0 {
					dialog.ShowError(errors.New("ДОСТИГНУТ РАЗМЕР 0"), myWindow)
					log.Println("Error: size 0")
					break
				}
			}
		} else if stepOptionSelect.Selected == STEP_OPTIONS[2] {
			h := minHeight
			w := minWidth
			for i := 0; i < int(cnt); i++ {
				points := f(xCenter, yCenter, w, h)
				pixels = getPixels(points)
				w += step
				h += step
				if h <= 0 || w <= 0 {
					dialog.ShowError(errors.New("ДОСТИГНУТ РАЗМЕР 0"), myWindow)
					log.Println("Error: size 0")
					break
				}
			}
		}
		raster.Refresh()
	})

	ellipseRangeCR := container.NewVBox(stepOptionLabel, stepOptionSelect, ellipseStepLabel, ellipseStepEntry)

	ellipseRangeCU := container.NewHBox(ellipseRangeCL, ellipseRangeCR)

	ellipseRangeC := container.NewVBox(ellipseRangeCU, ellipseCntLabel, ellipseCntEntry, ellipseRangeButton)

	sepRect5 := createHoriSepRect(circleC.MinSize().Width)

	timeMeasureOptionLabel1 := canvas.NewText("Сравнение временных", theme.ForegroundColor())
	timeMeasureOptionLabel1.TextSize = 14
	timeMeasureOptionLabel1.Alignment = fyne.TextAlignCenter
	timeMeasureOptionLabel2 := canvas.NewText("характеристик алгоритмов", theme.ForegroundColor())
	timeMeasureOptionLabel2.TextSize = 14
	timeMeasureOptionLabel2.Alignment = fyne.TextAlignCenter

	timeMeasureOptionSelect := widget.NewSelect(TIME_MEASURE_OPTIONS, func(value string) {
		log.Println("Select set to", value)
	})
	timeMeasureOptionSelect.SetSelected(TIME_MEASURE_OPTIONS[0])

	timeButton := widget.NewButton("Сравнить", func() {
		if timeMeasureOptionSelect.Selected == TIME_MEASURE_OPTIONS[0] {
			minRad, err := strconv.ParseFloat(circleMinRadEntry.Text, 64)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MIN R"), myWindow)
				log.Println("Error:", err)
				return
			}
			if minRad <= 0 {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MIN R"), myWindow)
				log.Println("Error: minRad <= 0")
				return
			}
			maxRad, err := strconv.ParseFloat(circleMaxRadEntry.Text, 64)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MAX R"), myWindow)
				log.Println("Error:", err)
				return
			}
			if maxRad <= 0 {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ MAX R"), myWindow)
				log.Println("Error: maxRad <= 0")
				return
			}
			step, err := strconv.ParseFloat(circleStepEntry.Text, 64)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
				log.Println("Error:", err)
				return
			}
			if (step > 0 && minRad > maxRad) || (step < 0 && minRad < maxRad) || step == 0 {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
				log.Println("Error: inf loop")
				return
			}
			fileName := MeasureCircle(minRad, maxRad, step)
			if fileName != "" {
				pic := myApp.NewWindow(fileName)
				image := canvas.NewImageFromFile(fileName)
				image.FillMode = canvas.ImageFillOriginal
				pic.SetContent(image)
				pic.Show()
			}
		} else {
			minHeight, err := strconv.ParseFloat(ellipseMinHeightEntry.Text, 64)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ВЫСОТА"), myWindow)
				log.Println("Error:", err)
				return
			}
			if minHeight <= 0 {
				dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ВЫСОТА"), myWindow)
				log.Println("Error: minHeight <= 0")
				return
			}
			minWidth, err := strconv.ParseFloat(ellipseMinWidthEntry.Text, 64)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ШИРИНА"), myWindow)
				log.Println("Error:", err)
				return
			}
			if minWidth <= 0 {
				dialog.ShowError(errors.New("НЕКОРРЕКТНАЯ MIN ШИРИНА"), myWindow)
				log.Println("Error: minWidth <= 0")
				return
			}
			step, err := strconv.ParseFloat(ellipseStepEntry.Text, 64)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНЫЙ ШАГ"), myWindow)
				log.Println("Error:", err)
				return
			}
			cnt, err := strconv.ParseInt(ellipseCntEntry.Text, 10, 32)
			if err != nil {
				dialog.ShowError(errors.New("НЕКОРРЕКТНОЕ КОЛИЧЕСТВО"), myWindow)
				log.Println("Error:", err)
				return
			}
			if cnt <= 0 {
				dialog.ShowError(errors.New("НЕКОРРЕКТНОЕ КОЛИЧЕСТВО"), myWindow)
				log.Println("Error: cnt <= 0")
				return
			}
			if math.Min(minHeight, minWidth)+step*float64(cnt) <= 0 {
				cnt -= int64(float64(cnt)+math.Min(minHeight, minWidth)/step) + 1
				dialog.ShowError(errors.New("БЫЛ ДОСТИГНУТ РАЗМЕР 0"), myWindow)
				log.Printf("Error: cnt = %d\n", cnt)
			}
			fileName := MeasureEllipse(minHeight, minWidth, step, cnt)
			if fileName != "" {
				pic := myApp.NewWindow(fileName)
				image := canvas.NewImageFromFile(fileName)
				image.FillMode = canvas.ImageFillOriginal
				pic.SetContent(image)
				pic.Show()
			}
		}
	})

	sepRect6 := createHoriSepRect(circleC.MinSize().Width)

	clearButton := widget.NewButton("Очистить экран", clearCanvas)

	copyrightLabel := canvas.NewText("Поляков А.И. ИУ7-42Б", theme.PressedColor())
	copyrightLabel.TextSize = 10
	copyrightLabel.Alignment = fyne.TextAlignTrailing

	upperFrame := container.NewHBox(methodC, colorC, sepRect1, centerC, sepRect2, circleC, sepRect3, ellipseC)

	raster = canvas.NewRasterWithPixels(drawCanvas)
	raster.SetMinSize(fyne.NewSize(upperFrame.MinSize().Width, 500))

	lowerFrame := container.NewPadded(raster)

	sepRectULWFrames := createHoriSepRect(upperFrame.MinSize().Width)

	combLFrame := container.NewVBox(upperFrame, sepRectULWFrames, lowerFrame)

	sepRectLRFrames := createVertSepRect(combLFrame.MinSize().Height)

	rightFrame := container.NewVBox(circleRangeC, sepRect4, ellipseRangeC, sepRect5,
		timeMeasureOptionLabel1, timeMeasureOptionLabel2, timeMeasureOptionSelect, timeButton, sepRect6, clearButton, copyrightLabel)

	content := container.NewHBox(combLFrame, sepRectLRFrames, rightFrame)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
