package main

import (
	"time"

	"github.com/seredot/trash/style"
)

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
			color:  style.ColorBullet,
		},
	}

	return b
}

func (b *Bullet) hit() {
	b.Object.speed = 0
	b.Object.sprite = '✧'
	b.Object.removeIn(time.Millisecond * 150)
}

func (b *Bullet) move() {
	b.Object.move()

	if b.x < float64(b.game.left) {
		b.x = float64(b.game.left)
		b.hit()
	}
	if b.x > float64(b.game.right)+0.5 {
		b.x = float64(b.game.right)
		b.hit()
	}
	if b.y < float64(b.game.top) {
		b.y = 0
		b.hit()
	}
	if b.y > float64(b.game.bottom)+0.5 {
		b.y = float64(b.game.bottom + 1)
		b.hit()
	}
}
