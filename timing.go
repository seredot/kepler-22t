package main

import "time"

type Timing interface {
	Frame() int
	DeltaT() time.Duration
}

func (g *Game) Frame() int {
	return g.frame
}

func (g *Game) DeltaT() time.Duration {
	return g.simuDelta
}
