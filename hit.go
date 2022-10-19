package main

import (
	"math"
	"time"

	"github.com/seredot/kepler-22t/color"
)

const HitRange = 0.75

func hitBullets(g *Game, bullets []*Bullet, aliens []*Alien) {
	for _, b := range bullets {
		for _, a := range aliens {
			if a.energy > 0 && !b.hasHit && math.Abs(b.x-a.x) < HitRange && math.Abs(b.y-a.y) < HitRange {
				a.energy = math.Max(0, a.energy-b.damage)
				a.fgColor = color.ColorAlien.Blend(color.Color{R: 0, G: 0, B: 0, A: 0.7 * (1 - a.energy/a.maxEnergy)})
				b.hit()
				g.addEffects(NewRedSpill(a.x, a.y))

				if a.energy <= 0 {
					g.score++
					a.sprite = '☠'
					a.speed = 0
					a.removeIn(time.Second)
				}
			}
		}
	}
}
