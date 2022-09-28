package main

import "github.com/gdamore/tcell/v2"

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
			color:  tcell.ColorBlue,
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

func (p *Player) draw() {
	p.Object.draw()

	for _, b := range p.bullets {
		b.draw()
	}
}

func (p *Player) fire() {
	p.bullets = append(p.bullets, NewBullet(p.game, p.x, p.y, p.dx, p.dy, 50))
}
