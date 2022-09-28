package main

import "math"

type Player struct {
	game   *Game
	x, y   float64
	dx, dy float64
	speed  float64
	drag   float64
	sprite rune
}

func NewPlayer(game *Game, x, y int) *Player {
	p := &Player{
		game:   game,
		x:      float64(x) - 1,
		y:      float64(y),
		dx:     1,
		dy:     0,
		sprite: 'X',
	}

	p.direction(p.dx, p.dy)
	return p
}

func (p *Player) direction(dx, dy float64) {
	p.speed = 10
	p.drag = 20

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
	p.x += p.dx * float64(p.game.deltaT) / 1000.0 * p.speed
	p.y += p.dy * float64(p.game.deltaT) / 1000.0 * p.speed
	p.speed = math.Max(0, p.speed-float64(p.game.deltaT)/1000.0*p.drag)

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

func (p *Player) scrX() int {
	return int(p.x)
}

func (p *Player) scrY() int {
	return int(p.y)
}

func (p *Player) draw() {
	p.move()
	p.game.screen.SetContent(p.scrX(), p.scrY(), p.sprite, nil, p.game.defStyle)
}
