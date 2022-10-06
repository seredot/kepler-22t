package vector

import "math"

func Mag(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func Norm(x, y float64) (nx, ny float64) {
	m := Mag(x, y)
	return x / m, y / m
}
