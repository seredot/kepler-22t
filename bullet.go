package main

import (
	"time"

	"github.com/seredot/kepler-22t/color"
)

type Bullet struct {
	Object
	damage float64
	hasHit bool
}

func NewBullet(x, y, dx, dy, speed float64) *Bullet {
	b := &Bullet{
		Object: Object{
			x:       float64(x),
			y:       float64(y),
			dx:      dx,
			dy:      dy,
			speed:   speed,
			sprite:  '•',
			fgColor: color.ColorBullet,
		},
		damage: 40.0,
		hasHit: false,
	}

	return b
}

func (b *Bullet) hit() {
	b.hasHit = true
	b.damage = 0
	b.Object.speed = 0
	b.Object.sprite = '✧'
	b.Object.removeIn(time.Millisecond * 150)
}

func (b *Bullet) move(t Timing, c Coords) {
	b.Object.move(t, c)

	if b.x < float64(c.Left()) {
		b.x = float64(c.Left())
		b.hit()
	}
	if b.x > float64(c.Right())+0.5 {
		b.x = float64(c.Right())
		b.hit()
	}
	if b.y < float64(c.Top()) {
		b.y = float64(c.Top())
		b.hit()
	}
	if b.y > float64(c.Bottom())+0.5 {
		b.y = float64(c.Bottom() + 1)
		b.hit()
	}
}
