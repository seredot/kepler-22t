package main

import "time"

type Gun interface {
	Name() string
	Delay() time.Duration
}

type Pistol struct{}

func (g Pistol) Name() string {
	return "Pistol"
}

func (g Pistol) Delay() time.Duration {
	return 300 * time.Millisecond
}

type MachineGun struct{}

func (g MachineGun) Name() string {
	return "Machine Gun"
}

func (g MachineGun) Delay() time.Duration {
	return 100 * time.Millisecond
}
