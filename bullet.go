package main

type Bullet struct {
	Object
	sprite rune
}

func NewBullet(game *Game, x, y, dx, dy, speed float64) *Bullet {
	b := &Bullet{
		Object: Object{
			game: game,
			x:    float64(x) - 1,
			y:    float64(y),
			dx:   dx,
			dy:   dy,
		},
		sprite: '•',
	}

	return b
}

func (b *Bullet) draw() {
	b.move()
	b.game.screen.SetContent(b.scrX(), b.scrY(), b.sprite, nil, b.game.defStyle)
}
