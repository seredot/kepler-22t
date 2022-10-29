package game

import "github.com/seredot/kepler-22t/color"

// drawFog draws an animating fog effect on the whole screen using Perlin Noise
// like algorithm. The 2D render actually passes through a 3D noise outputting
// slices of Z axis over time.
func (g *Game) drawFog() {
	g.ResetStyle()
	g.Foreground(color.ColorTransparent)

	zoom := 8.0
	ambient := 0.15
	strength := 0.3
	seconds := float64(g.totalT.Seconds()) * .25

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			alpha := ambient + strength*g.noise.Eval3(
				float64(x)/zoom+seconds*0.4,
				float64(y)/zoom+seconds*0.4,
				seconds,
			)
			c := color.Color{R: .6, G: 1, B: .6, A: alpha}
			g.Background(c)

			g.PutColor(x, y)
		}
	}

	g.ResetStyle()
}
