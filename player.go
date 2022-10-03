package main

import (
	"math"

	"github.com/seredot/trash/style"
)

type Player struct {
	Object
	bullets []*Bullet
}

func NewPlayer(game *Game, x, y int) *Player {
	p := &Player{
		Object: Object{
			game:   game,
			x:      float64(x) - 1,
			y:      float64(y),
			dx:     1,
			dy:     0,
			drag:   20,
			sprite: 'X',
			color:  style.ColorPlayer,
		},
		bullets: []*Bullet{},
	}

	p.direction(p.dx, p.dy)
	return p
}

func (p *Player) direction(dx, dy float64) {
	p.speed = 10

	p.dx = dx
	p.dy = dy

	if dx == -1 {
		p.sprite = '◀'
	} else if dx == 1 {
		p.sprite = '▶'
	} else if dy == -1 {
		p.sprite = '▲'
	} else {
		p.sprite = '▼'
	}
}

func (p *Player) move() {
	p.Object.move()

	if p.x < 1 {
		p.x = 1
		p.speed = 0
	}
	if p.x > float64(p.game.width)-2 {
		p.x = float64(p.game.width) - 2
		p.speed = 0
	}
	if p.y < 1 {
		p.y = 1
		p.speed = 0
	}
	if p.y > float64(p.game.height)-2 {
		p.y = float64(p.game.height) - 2
		p.speed = 0
	}
}

func (p *Player) draw() {
	p.Object.draw()
}

func (p *Player) fire() {
	dx := float64(p.game.mouseX) - p.x
	dy := float64(p.game.mouseY) - p.y
	mag := math.Sqrt(dx*dx + dy*dy)
	dx /= mag
	dy /= mag
	p.bullets = append(p.bullets, NewBullet(p.game, p.x, p.y, dx, dy, 30))
}
