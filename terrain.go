package main

import "github.com/seredot/kepler-22t/color"

func (g *Game) drawTerrain() {
	g.ResetStyle()

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			zoom := 8.0
			ambient := 10.0
			strength := 12.0
			luminosity := ambient + strength*g.noise.Eval3(float64(x)/zoom, float64(y)/zoom, float64(g.totalT.Seconds()))
			g.Background(color.Hsl2Rgb(242, 26, luminosity))
			g.PutChar(x, y, ' ')
		}
	}

	g.ResetStyle()
}
