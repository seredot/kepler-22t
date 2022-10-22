package gun

import (
	"time"

	"github.com/seredot/kepler-22t/object"
	"github.com/seredot/kepler-22t/screen"
)

type Pistol struct{ gun }

func NewPistol() Gun {
	return &machineGun{
		gun: gun{
			name:  "Pistol",
			delay: 300 * time.Millisecond,
		},
	}
}

func (g Pistol) Fire(c screen.Coords) []*object.Bullet {
	dx, dy := fireVector(c)

	return []*object.Bullet{object.NewBullet(c.PlayerX(), c.PlayerY(), dx, dy, 30.0)}
}
