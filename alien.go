package main

import (
	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/vector"
)

type Alien struct {
	Object
	maxEnergy float64
	energy    float64
	damage    float64 // damage per second
	reaches   bool    // player is in reach and getting damage
}

func NewAlien(game *Game) *Alien {
	x := float64(game.left) + float64(game.right-game.left)*game.noise.Eval2(0, float64(game.frame))
	y := float64(game.top) + float64(game.bottom-game.top)*game.noise.Eval2(float64(game.frame), 0)

	e := &Alien{
		Object: Object{
			x:       x,
			y:       y,
			speed:   2.0,
			sprite:  '✹',
			fgColor: color.ColorAlien,
		},
		maxEnergy: 100.0,
		energy:    100.0,
		damage:    5.0,
	}

	return e
}

func (a *Alien) move(t Timing, c Coords) {
	dx := c.PlayerX() - a.x
	dy := c.PlayerY() - a.y
	dist := vector.Mag(dx, dy)
	a.dx, a.dy = vector.Norm(dx, dy)

	reaches := false

	if dist > 1.0 && a.energy > 0 {
		a.Object.move(t, c)
	}

	if dist <= 1.0 && a.energy > 0 {
		reaches = true
	}

	a.reaches = reaches
}
