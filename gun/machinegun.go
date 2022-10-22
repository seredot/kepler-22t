package gun

import (
	"time"

	"github.com/seredot/kepler-22t/object"
	"github.com/seredot/kepler-22t/screen"
)

type machineGun struct{ gun }

func NewMachineGun() Gun {
	return &machineGun{
		gun: gun{
			name:  "Machine Gun",
			delay: 100 * time.Millisecond,
		},
	}
}

func (g machineGun) Fire(c screen.Coords) []*object.Bullet {
	dx, dy := fireVector(c)

	return []*object.Bullet{object.NewBullet(c.PlayerX(), c.PlayerY(), dx, dy, 60.0)}
}
