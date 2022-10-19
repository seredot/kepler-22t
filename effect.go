package main

import (
	"time"

	"github.com/seredot/kepler-22t/color"
)

type Effect struct {
	Object

	fromFg   color.Color
	toFg     color.Color
	fromBg   color.Color
	toBg     color.Color
	fromTime time.Time
	toTime   time.Time
}

func (e *Effect) move(t Timing, c Coords) {
	now := time.Now()
	elapsed := float64(now.Sub(e.fromTime))
	total := float64(e.toTime.Sub(e.fromTime))

	if total <= 0 || now.After(e.toTime) {
		e.removed = true
		return
	}

	ratio := elapsed / total

	e.fgColor = e.fromFg.Interpolate(e.toFg, ratio)
	e.bgColor = e.fromBg.Interpolate(e.toBg, ratio)
}

func (e *Effect) draw(c Canvas) {
	c.ResetStyle()
	c.Foreground(e.fgColor)
	c.Background(e.bgColor)
	c.PutColor(e.scrX(), e.scrY())
	c.ResetStyle()
}

func NewRedSpill(x, y float64) *Effect {
	return &Effect{
		Object: Object{
			x: x,
			y: y,
		},
		fromFg:   color.ColorTransparent,
		toFg:     color.ColorTransparent,
		fromBg:   color.ColorRedSpill,
		toBg:     color.ColorTransparent,
		fromTime: time.Now(),
		toTime:   time.Now().Add(5 * time.Second),
	}
}
