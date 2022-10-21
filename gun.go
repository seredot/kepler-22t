package main

import (
	"time"

	"github.com/seredot/kepler-22t/vector"
)

type Gun interface {
	Name() string
	Delay() time.Duration
	Fire(c Coords) []*Bullet
}

type Pistol struct{}

func (g Pistol) Name() string {
	return "Pistol"
}

func (g Pistol) Delay() time.Duration {
	return 300 * time.Millisecond
}

func (g Pistol) Fire(c Coords) []*Bullet {
	dx, dy := fireVector(c)

	return []*Bullet{NewBullet(c.PlayerX(), c.PlayerY(), dx, dy, 30.0)}
}

type MachineGun struct{}

func (g MachineGun) Name() string {
	return "Machine Gun"
}

func (g MachineGun) Delay() time.Duration {
	return 100 * time.Millisecond
}

func (g MachineGun) Fire(c Coords) []*Bullet {
	dx, dy := fireVector(c)

	return []*Bullet{NewBullet(c.PlayerX(), c.PlayerY(), dx, dy, 60.0)}
}

func fireVector(c Coords) (dx, dy float64) {
	dx = float64(c.MouseX()) - c.PlayerX()
	dy = float64(c.MouseY()) - c.PlayerY()
	dx, dy = vector.Norm(dx, dy)

	return
}
