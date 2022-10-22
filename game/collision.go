package game

import (
	"math"

	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/object"
)

const HitRange = 0.75

func (g *Game) checkCollisions() {
	for _, a := range g.aliens {
		// Check if alien reaches the player
		if a.Reaches {
			g.health = math.Max(0, g.health-a.Damage*g.coords.DeltaT().Seconds())

			if g.health == 0 {
				g.state = GameOver
			}
		}

		// Check if bullets hit the alien
		for _, b := range g.bullets {
			if a.Energy > 0 && !b.HasHit() && math.Abs(b.X-a.X) < HitRange && math.Abs(b.Y-a.Y) < HitRange {
				a.GetDamage(b.Damage())
				a.FgColor = color.ColorAlien.Blend(color.Color{R: 0, G: 0, B: 0, A: 0.7 * (1 - a.Energy/a.MaxEnergy)})
				b.Hit()
				g.addEffects(object.NewRedSpill(a.X, a.Y))

				if a.Dead {
					g.score++
				}
			}
		}
	}
}
