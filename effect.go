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

func NewGunFlash(x, y float64) []*Effect {
	l0 := Effect{
		Object: Object{
			x: x,
			y: y,
		},
		fromFg:   color.ColorTransparent,
		toFg:     color.ColorTransparent,
		fromBg:   color.ColorBullet,
		toBg:     color.ColorTransparent,
		fromTime: time.Now(),
		toTime:   time.Now().Add(100 * time.Millisecond),
	}

	l0.fromBg.A = 0.3
	l1 := l0
	l1.x++
	l1.fromBg.A = 0.15
	l2 := l0
	l2.y++
	l2.fromBg.A = 0.1
	l3 := l0
	l3.x--
	l3.fromBg.A = 0.15
	l4 := l0
	l4.y--
	l4.fromBg.A = 0.1
	l5 := l0
	l5.x++
	l5.y++
	l5.fromBg.A = 0.05
	l6 := l0
	l6.x--
	l6.y++
	l6.fromBg.A = 0.05
	l7 := l0
	l7.x--
	l7.y--
	l7.fromBg.A = 0.05
	l8 := l0
	l8.x++
	l8.y--
	l8.fromBg.A = 0.05

	return []*Effect{&l0, &l1, &l2, &l3, &l4, &l5, &l6, &l7, &l8}
}
