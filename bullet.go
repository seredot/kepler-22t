package main

import "github.com/gdamore/tcell/v2"

type Bullet struct {
	Object
}

func NewBullet(game *Game, x, y, dx, dy, speed float64) *Bullet {
	b := &Bullet{
		Object: Object{
			game:   game,
			x:      float64(x),
			y:      float64(y),
			dx:     dx,
			dy:     dy,
			speed:  speed,
			sprite: '•',
			color:  tcell.ColorYellow,
		},
	}

	return b
}
