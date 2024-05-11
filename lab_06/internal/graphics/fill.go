package graphics

import (
	"github.com/golang-collections/collections/stack"
)

func pixelFilled(pixels *SafePixels, p IPoint) bool {
	_, keyExists := pixels.PXS[p]
	return keyExists
}

func fillPixel(pixels *SafePixels, p IPoint) {
	(*pixels).MU.Lock()
	(*pixels).PXS[p] = true
	(*pixels).MU.Unlock()
}

func findPixel(stack *stack.Stack, pixels *SafePixels, xRight, x, y int) {
	p := IPoint{X: x, Y: y}
	for p.X <= xRight {
		flag := false
		for !pixelFilled(pixels, p) && p.X <= xRight {
			if !flag {
				flag = true
			}
			p.X++
		}

		if flag {
			if p.X == xRight && !pixelFilled(pixels, p) {
				stack.Push(p)
			} else {
				stack.Push(IPoint{p.X - 1, p.Y})
			}
			flag = false
		}

		xTemp := p.X
		for pixelFilled(pixels, p) && p.X < xRight {
			p.X++
		}

		if p.X == xTemp {
			p.X++
		}
	}
}

func fillRight(pixels *SafePixels, p IPoint) int {
	for !pixelFilled(pixels, p) {
		fillPixel(pixels, p)
		p.X++
	}
	return int(p.X - 1)
}

func fillLeft(pixels *SafePixels, p IPoint) int {
	for !pixelFilled(pixels, p) {
		fillPixel(pixels, p)
		p.X--
	}
	return int(p.X + 1)
}

func Fill(pixels *SafePixels, fillPoint FPoint) {
	stack := stack.New()

	stack.Push(ftoiPoint(fillPoint))

	for stack.Len() > 0 {
		point := stack.Pop().(IPoint)
		fillPixel(pixels, point)
		xRight := fillRight(pixels, IPoint{X: point.X + 1, Y: point.Y})
		xLeft := fillLeft(pixels, IPoint{X: point.X - 1, Y: point.Y})
		findPixel(stack, pixels, xRight, xLeft, point.Y+1)
		findPixel(stack, pixels, xRight, xLeft, point.Y-1)
	}
}

func FillWDelay(pixels *SafePixels, fillPoint FPoint, c chan int) {
	stack := stack.New()

	stack.Push(ftoiPoint(fillPoint))

	for stack.Len() > 0 {
		point := stack.Pop().(IPoint)
		fillPixel(pixels, point)
		xRight := fillRight(pixels, IPoint{X: point.X + 1, Y: point.Y})
		xLeft := fillLeft(pixels, IPoint{X: point.X - 1, Y: point.Y})
		findPixel(stack, pixels, xRight, xLeft, point.Y+1)
		findPixel(stack, pixels, xRight, xLeft, point.Y-1)
		c <- 1
	}
	close(c)
}
