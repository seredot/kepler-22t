package object

import (
	"time"

	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/screen"
)

type Bullet struct {
	Object
	damage float64
	hasHit bool
}

func NewBullet(x, y, dx, dy, speed float64) *Bullet {
	b := &Bullet{
		Object: NewObject(
			float64(x), float64(y),
			dx,
			dy,
			speed,
			0,
			'•',
			color.ColorAmber,
			color.ColorTransparent,
		),
		damage: 40.0,
		hasHit: false,
	}

	return b
}

func (b *Bullet) HasHit() bool {
	return b.hasHit
}

func (b *Bullet) Hit() {
	b.hasHit = true
	b.damage = 0
	b.Speed = 0
	b.Sprite = '✧'
	b.removeIn(time.Millisecond * 150)
}

func (b *Bullet) Damage() float64 {
	return b.damage
}

func (b *Bullet) Move(c screen.Coords) {
	b.Object.Move(c)

	if b.X < float64(c.Left()) {
		b.X = float64(c.Left())
		b.Hit()
	}
	if b.X > float64(c.Right())+0.5 {
		b.X = float64(c.Right())
		b.Hit()
	}
	if b.Y < float64(c.Top()) {
		b.Y = float64(c.Top())
		b.Hit()
	}
	if b.Y > float64(c.Bottom())+0.5 {
		b.Y = float64(c.Bottom() + 1)
		b.Hit()
	}
}
