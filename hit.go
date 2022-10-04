package main

import (
	"math"

	"github.com/seredot/trash/style"
)

const HitRange = 0.75

func hitBullets(bullets []*Bullet, enemies []*Enemy) {
	for _, b := range bullets {
		for _, e := range enemies {
			if math.Abs(b.x-e.x) < HitRange && math.Abs(b.y-e.y) < HitRange {
				e.color = style.ColorPointer
			}
		}
	}
}
