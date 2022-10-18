package main

import "github.com/seredot/kepler-22t/color"

func (g *Game) drawFog() {
	g.ResetStyle()

	zoom := 8.0
	ambient := 0.15
	strength := 0.2

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			luminosity := ambient + strength*g.noise.Eval3(float64(x)/zoom, float64(y)/zoom, float64(g.totalT.Seconds()))
			c := color.Color{R: .6, G: .6, B: 1, A: luminosity}
			g.Background(c)
			g.Foreground(c)

			g.PutColor(x, y)
		}
	}

	g.ResetStyle()
}
