package main

import (
	"math"
	"time"

	"github.com/seredot/kepler-22t/color"
)

const HitRange = 0.75

func hitBullets(g *Game, bullets []*Bullet, aliens []*Alien) {
	for _, b := range bullets {
		for _, e := range aliens {
			if e.energy > 0 && !b.hasHit && math.Abs(b.x-e.x) < HitRange && math.Abs(b.y-e.y) < HitRange {
				e.energy = math.Max(0, e.energy-b.damage)
				e.fgColor = color.ColorAlien.Blend(color.Color{R: 0, G: 0, B: 0, A: 0.7 * (1 - e.energy/e.maxEnergy)})
				b.hit()
				g.addEffects(NewRedSpill(e.x, e.y))

				if e.energy <= 0 {
					g.score++
					e.speed = 0
					e.removeIn(time.Second)
				}
			}
		}
	}
}
