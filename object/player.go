package object

import (
	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/screen"
)

type Player struct {
	Object
}

func NewPlayer(x, y int) *Player {
	p := &Player{
		Object: Object{
			X:       float64(x) - 1,
			Y:       float64(y),
			Dx:      1,
			Dy:      0,
			Drag:    20,
			Sprite:  '◉',
			FgColor: color.ColorPlayer,
		},
	}

	p.Direction(p.Dx, p.Dy)
	return p
}

func (p *Player) Direction(dx, dy float64) {
	p.Speed = 10

	p.Dx = dx
	p.Dy = dy
}

func (p *Player) Move(c screen.Coords) {
	p.Object.Move(c)

	if p.X < float64(c.Left()) {
		p.X = float64(c.Left())
		p.Speed = 0
	}
	if p.X > float64(c.Right()) {
		p.X = float64(c.Right())
		p.Speed = 0
	}
	if p.Y < float64(c.Top()) {
		p.Y = float64(c.Top())
		p.Speed = 0
	}
	if p.Y > float64(c.Bottom()) {
		p.Y = float64(c.Bottom())
		p.Speed = 0
	}
}

func (p *Player) Draw(c screen.Canvas) {
	p.Object.Draw(c)
}
