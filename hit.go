package main

import (
	"math"

	"github.com/seredot/kepler-22t/style"
)

const HitRange = 0.75

func hitBullets(bullets []*Bullet, aliens []*Alien) {
	for _, b := range bullets {
		for _, e := range aliens {
			if math.Abs(b.x-e.x) < HitRange && math.Abs(b.y-e.y) < HitRange {
				e.color = style.ColorPointer
			}
		}
	}
}
