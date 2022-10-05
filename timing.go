package main

type Timing interface {
	Frame() int
	DeltaT() float64
}

func (g *Game) Frame() int {
	return g.frame
}

func (g *Game) DeltaT() float64 {
	return float64(g.deltaT)
}
