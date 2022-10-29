package object

import (
	"time"

	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/screen"
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

func (e *Effect) Move(c screen.Coords) {
	now := time.Now()
	elapsed := float64(now.Sub(e.fromTime))
	total := float64(e.toTime.Sub(e.fromTime))

	if total <= 0 || now.After(e.toTime) {
		e.Removed = true
		return
	}

	ratio := elapsed / total

	e.FgColor = e.fromFg.Interpolate(e.toFg, ratio)
	e.BgColor = e.fromBg.Interpolate(e.toBg, ratio)
}

func (e *Effect) Draw(c screen.Canvas) {
	c.ResetStyle()
	c.Foreground(e.FgColor)
	c.Background(e.BgColor)
	c.PutColor(e.ScrX(), e.ScrY())
	c.ResetStyle()
}

func NewRedSpill(x, y float64) *Effect {
	return &Effect{
		Object: Object{
			X: x,
			Y: y,
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
			X: x,
			Y: y,
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
	l1.X++
	l1.fromBg.A = 0.15
	l2 := l0
	l2.Y++
	l2.fromBg.A = 0.1
	l3 := l0
	l3.X--
	l3.fromBg.A = 0.15
	l4 := l0
	l4.Y--
	l4.fromBg.A = 0.1
	l5 := l0
	l5.X++
	l5.Y++
	l5.fromBg.A = 0.05
	l6 := l0
	l6.X--
	l6.Y++
	l6.fromBg.A = 0.05
	l7 := l0
	l7.X--
	l7.Y--
	l7.fromBg.A = 0.05
	l8 := l0
	l8.X++
	l8.Y--
	l8.fromBg.A = 0.05

	return []*Effect{&l0, &l1, &l2, &l3, &l4, &l5, &l6, &l7, &l8}
}
