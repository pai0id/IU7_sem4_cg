package graphics

import "math"

var CurrFunc SurfaceEq

type SurfaceEq func(x, z float64) float64

var FuncArr = []SurfaceEq{
	func(x, z float64) float64 { return 0 },
	func(x, z float64) float64 { return math.Sin(x) * math.Sin(z) },
	func(x, z float64) float64 { return math.Sqrt(x*x + z*z) },
	func(x, z float64) float64 { return x*x + z*z },
}

var FuncStrArr = []string{
	"0",
	"sin(x) * sin(z)",
	"sqrt(x^2 + z^2)",
	"x^2 + z^2",
}

func GetID(str string) SurfaceEq {
	for i, s := range FuncStrArr {
		if s == str {
			return FuncArr[i]
		}
	}
	return nil
}
