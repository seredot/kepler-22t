package color

import (
	"math"
)

type Color struct {
	R, G, B, A float64
}

var (
	ColorTransparent = Color{0, 0, 0, 0}
	ColorBlack       = Color{0, 0, 0, 1}
	ColorWhite       = Color{.9, .9, .9, 1}
	ColorGrey        = Color{.4, .4, .4, 1}
	ColorAmber       = Color{1, .75, 0, 1}
	ColorCrossRed    = Color{.8, 0, 0, 1}
	ColorPlayer      = Color{.9, .9, 1, 1}
	ColorPointer     = Color{0, 1, 0, 1}
	ColorBullet      = Color{1, 1, 0, 1}
	ColorAlien       = Color{0, .7, .1, 1}
	ColorRedSpill    = Color{1, 0, 0, 0.2}
	ColorBox         = Color{.259, .475, .282, 1}
)

func NewColorIntRGBA(r, g, b, a uint64) Color {
	return Color{
		R: float64(r) / 255.0,
		G: float64(g) / 255.0,
		B: float64(b) / 255.0,
		A: float64(a) / 255.0,
	}
}

func NewColorIntRGB(r, g, b uint64) Color {
	return NewColorIntRGBA(r, g, b, 255)
}

func (c Color) Blend(a Color) Color {
	return Color{
		R: c.R*(1-a.A) + a.R*a.A,
		G: c.G*(1-a.A) + a.G*a.A,
		B: c.B*(1-a.A) + a.B*a.A,
		A: 1,
	}
}

func (c Color) Interpolate(a Color, r float64) Color {
	return Color{
		R: c.R*(1-r) + a.R*r,
		G: c.G*(1-r) + a.G*r,
		B: c.B*(1-r) + a.B*r,
		A: c.A*(1-r) + a.A*r,
	}
}

// Function is based on https://github.com/hisamafahri/coco/blob/main/hsl.go
func Hsl2Rgb(h float64, s float64, l float64) Color {
	h = h / 360
	s = s / 100
	l = l / 100

	var t2 float64
	var t3 float64
	var val float64
	var result [3]uint64

	if s == 0 {
		val = l * 255
		result[0] = uint64(math.Round(val))
		result[1] = uint64(math.Round(val))
		result[2] = uint64(math.Round(val))

		return NewColorIntRGB(result[0], result[1], result[2])
	}

	if l < 0.5 {
		t2 = l * (1 + s)
	} else {
		t2 = l + s - l*s
	}

	t1 := 2*l - t2

	rgb := [3]float64{0, 0, 0}

	for i := 0; i < 3; i++ {
		t3 = h + 1.0/3.0*(-(float64(i) - 1))

		if t3 < 0 {
			t3++
		}

		if t3 > 1 {
			t3--
		}

		if (6 * t3) < 1 {
			val = t1 + (t2-t1)*6*t3
		} else if (2 * t3) < 1 {
			val = t2
		} else if (3 * t3) < 2 {
			val = t1 + (t2-t1)*(2.0/3.0-t3)*6
		} else {
			val = t1
		}
		rgb[i] = val * 255
	}

	result[0] = uint64(math.Round(rgb[0]))
	result[1] = uint64(math.Round(rgb[1]))
	result[2] = uint64(math.Round(rgb[2]))

	return NewColorIntRGB(result[0], result[1], result[2])
}
