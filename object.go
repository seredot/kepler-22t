package main

import (
	"math"

	"github.com/seredot/trash/style"
)

type Object struct {
	game   *Game
	x, y   float64
	dx, dy float64
	speed  float64
	drag   float64
	sprite rune
	color  style.Color
}

func (o *Object) move() {
	o.x += o.dx * float64(o.game.deltaT) / 1000.0 * o.speed
	o.y += o.dy * float64(o.game.deltaT) / 1000.0 * o.speed
	o.speed = math.Max(0, o.speed-float64(o.game.deltaT)/1000.0*o.drag)

	if o.x < 1 {
		o.x = 1
		o.speed = 0
	}
	if o.x > float64(o.game.width)-2 {
		o.x = float64(o.game.width) - 2
		o.speed = 0
	}
	if o.y < 1 {
		o.y = 1
		o.speed = 0
	}
	if o.y > float64(o.game.height)-2 {
		o.y = float64(o.game.height) - 2
		o.speed = 0
	}
}

func (o *Object) scrX() int {
	return int(math.Round(o.x))
}

func (o *Object) scrY() int {
	return int(math.Round(o.y))
}

func (o *Object) draw() {
	o.move()
	o.game.ResetStyle()
	o.game.Foreground(o.color)
	o.game.PatchChar(o.scrX(), o.scrY(), o.sprite)
	o.game.ResetStyle()
}
