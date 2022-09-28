package main

import "math"

type Object struct {
	game   *Game
	x, y   float64
	dx, dy float64
	speed  float64
	drag   float64
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
	return int(o.x)
}

func (o *Object) scrY() int {
	return int(o.y)
}
