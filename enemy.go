package main

import (
	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/vector"
)

type Alien struct {
	Object
	energy float64
}

func NewAlien(game *Game) *Alien {
	x := float64(game.left) + float64(game.right-game.left)*game.noise.Eval2(0, float64(game.frame))
	y := float64(game.top) + float64(game.bottom-game.top)*game.noise.Eval2(float64(game.frame), 0)

	e := &Alien{
		Object: Object{
			x:      x,
			y:      y,
			speed:  1,
			sprite: '✹',
			color:  color.ColorAlien,
		},
		energy: 100.0,
	}

	return e
}

func (e *Alien) move(t Timing, c Coords) {
	dx := c.PlayerX() - e.x
	dy := c.PlayerY() - e.y
	dist := vector.Mag(dx, dy)
	e.dx, e.dy = vector.Norm(dx, dy)

	if dist > 1.0 {
		e.Object.move(t)
	}
}
