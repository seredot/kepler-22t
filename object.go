package main

import (
	"math"
	"time"

	"github.com/seredot/kepler-22t/color"
)

type Object struct {
	x, y    float64
	dx, dy  float64
	speed   float64
	drag    float64
	sprite  rune
	color   color.Color
	removed bool
}

func (o *Object) removeIn(t time.Duration) {
	time.AfterFunc(t, func() {
		o.removed = true
	})
}

func (o *Object) move(t Timing) {
	o.x += o.dx * t.DeltaT().Seconds() * o.speed
	o.y += o.dy * t.DeltaT().Seconds() * o.speed
	o.speed = math.Max(0, o.speed-t.DeltaT().Seconds()*o.drag)
}

func (o *Object) scrX() int {
	return int(math.Round(o.x))
}

func (o *Object) scrY() int {
	return int(math.Round(o.y))
}

func (o *Object) draw(c Canvas) {
	c.ResetStyle()
	c.Foreground(o.color)
	c.PatchChar(o.scrX(), o.scrY(), o.sprite)
	c.ResetStyle()
}
