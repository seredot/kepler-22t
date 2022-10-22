package object

import (
	"math"
	"time"

	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/screen"
)

type Object struct {
	X, Y    float64
	Dx, Dy  float64
	Speed   float64
	Drag    float64
	Sprite  rune
	FgColor color.Color
	BgColor color.Color
	Removed bool
}

func NewObject(
	x, y,
	dx, dy,
	speed,
	drag float64,
	sprite rune,
	fgColor color.Color,
	bgColor color.Color,
) Object {
	return Object{
		X:       float64(x),
		Y:       float64(y),
		Dx:      dx,
		Dy:      dy,
		Speed:   speed,
		Sprite:  '•',
		FgColor: color.ColorAmber,
	}
}

func (o *Object) removeIn(t time.Duration) {
	time.AfterFunc(t, func() {
		o.Removed = true
	})
}

func (o *Object) Move(c screen.Coords) {
	o.X += o.Dx * c.DeltaT().Seconds() * o.Speed
	o.Y += o.Dy * c.DeltaT().Seconds() * o.Speed
	o.Speed = math.Max(0, o.Speed-c.DeltaT().Seconds()*o.Drag)
}

func (o *Object) scrX() int {
	return int(math.Round(o.X))
}

func (o *Object) scrY() int {
	return int(math.Round(o.Y))
}

func (o *Object) Draw(c screen.Canvas) {
	c.ResetStyle()
	c.Foreground(o.FgColor)
	c.Background(o.BgColor)
	c.PatchChar(o.scrX(), o.scrY(), o.Sprite)
	c.ResetStyle()
}
