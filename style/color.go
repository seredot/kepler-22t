package style

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

func ColorRGB(r, g, b uint64) Color {
	return Color(uint64(tcell.ColorIsRGB) | uint64(tcell.ColorValid) | r<<16 | g<<8 | b)
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

		return ColorRGB(result[0], result[1], result[2])
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

	return ColorRGB(result[0], result[1], result[2])
}
