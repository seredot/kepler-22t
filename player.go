package main

import "github.com/seredot/kepler-22t/color"

type Player struct {
	Object
}

func NewPlayer(game *Game, x, y int) *Player {
	p := &Player{
		Object: Object{
			x:       float64(x) - 1,
			y:       float64(y),
			dx:      1,
			dy:      0,
			drag:    20,
			sprite:  '◉',
			fgColor: color.ColorPlayer,
		},
	}

	p.direction(p.dx, p.dy)
	return p
}

func (p *Player) direction(dx, dy float64) {
	p.speed = 10

	p.dx = dx
	p.dy = dy
}

func (p *Player) move(t Timing, c Coords) {
	p.Object.move(t, c)

	if p.x < float64(c.Left()) {
		p.x = float64(c.Left())
		p.speed = 0
	}
	if p.x > float64(c.Right()) {
		p.x = float64(c.Right())
		p.speed = 0
	}
	if p.y < float64(c.Top()) {
		p.y = float64(c.Top())
		p.speed = 0
	}
	if p.y > float64(c.Bottom()) {
		p.y = float64(c.Bottom())
		p.speed = 0
	}
}

func (p *Player) draw(c Canvas) {
	p.Object.draw(c)
}
