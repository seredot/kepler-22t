package main

type Player struct {
	Object
	sprite rune
}

func NewPlayer(game *Game, x, y int) *Player {
	p := &Player{
		Object: Object{
			game: game,
			x:    float64(x) - 1,
			y:    float64(y),
			dx:   1,
			dy:   0,
			drag: 20,
		},
		sprite: 'X',
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
	p.move()
	p.game.screen.SetContent(p.scrX(), p.scrY(), p.sprite, nil, p.game.defStyle)
}
