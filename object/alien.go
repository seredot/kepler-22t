package object

import (
	"math"
	"time"

	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/screen"
	"github.com/seredot/kepler-22t/vector"
)

type Alien struct {
	Object
	MaxEnergy float64
	Energy    float64
	Damage    float64 // damage per second
	Reaches   bool    // player is in reach and getting damage
	Dead      bool
}

func NewAlien(c screen.Canvas) *Alien {
	x := float64(c.Left()) + float64(c.Right()-c.Left())*c.Noise().Eval2(0, float64(c.Frame()))
	y := float64(c.Top()) + float64(c.Bottom()-c.Top())*c.Noise().Eval2(float64(c.Frame()), 0)

	e := &Alien{
		Object: Object{
			X:       x,
			Y:       y,
			Speed:   3.0,
			Sprite:  '⚉',
			FgColor: color.ColorAlien,
		},
		MaxEnergy: 100.0,
		Energy:    100.0,
		Damage:    5.0,
	}

	return e
}

func (a *Alien) Move(c screen.Coords) {
	dx := c.PlayerX() - a.X
	dy := c.PlayerY() - a.Y
	dist := vector.Mag(dx, dy)
	a.Dx, a.Dy = vector.Norm(dx, dy)

	reaches := false

	// Run away if wounded
	if a.Energy < 30 && a.Speed > 0 {
		a.Dx *= -1
		a.Dy *= -1
	}

	if dist > 1.0 && a.Energy > 0 {
		a.Object.Move(c)
	}

	if c.OutOfScreen(a.ScrX(), a.ScrY()) {
		a.removeIn(0)
	}

	if dist <= 1.0 && a.Energy > 0 {
		reaches = true
	}

	a.Reaches = reaches
}

func (a *Alien) GetDamage(d float64) {
	a.Energy = math.Max(0, a.Energy-d)

	if a.Energy < 50 {
		a.Sprite = '⚈'
	}

	if a.Energy <= 0 {
		a.die()
	}
}

func (a *Alien) die() {
	a.Dead = true
	a.FgColor = color.ColorBlack
	a.Sprite = '☠'
	a.Speed = 0
	a.removeIn(time.Second)
}
