package screen

import (
	"time"

	"github.com/ojrac/opensimplex-go"
	"github.com/seredot/kepler-22t/color"
)

type Coords interface {
	Width() int
	Height() int
	Left() int
	Right() int
	Top() int
	Bottom() int
	MouseX() int
	MouseY() int
	PlayerX() float64
	PlayerY() float64
	OutOfScreen(x, y int) bool
	Frame() int
	DeltaT() time.Duration
}

type Canvas interface {
	Coords

	ResetStyle()
	Background(c color.Color)
	Foreground(c color.Color)
	PutChar(x, y int, r rune)
	PatchChar(x, y int, r rune)
	PutColor(x, y int)
	DrawText(x, y int, text string)
	Noise() opensimplex.Noise
}
