package gun

import (
	"time"

	"github.com/seredot/kepler-22t/object"
	"github.com/seredot/kepler-22t/screen"
	"github.com/seredot/kepler-22t/vector"
)

type gun struct {
	name  string
	delay time.Duration
}

type Gun interface {
	Name() string
	Delay() time.Duration
	Fire(c screen.Coords) []*object.Bullet
}

func (g gun) Name() string {
	return g.name
}

func (g gun) Delay() time.Duration {
	return g.delay
}

func fireVector(c screen.Coords) (dx, dy float64) {
	dx = float64(c.MouseX()) - c.PlayerX()
	dy = float64(c.MouseY()) - c.PlayerY()
	dx, dy = vector.Norm(dx, dy)

	return
}
