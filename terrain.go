package main

import "github.com/seredot/trash/style"

func (g *Game) drawTerrain() {

	g.ResetStyle()
	for x := g.left; x <= g.right; x++ {
		for y := g.top; y <= g.bottom; y++ {
			zoom := 16.0
			ambient := 10.0
			strength := 12.0
			luminosity := ambient + strength*g.noise.Eval2(float64(x)/zoom, float64(y)/zoom)
			g.Background(style.Color(style.Hsl2Rgb(242, 26, luminosity)))
			g.PutChar(x, y, ' ')
		}
	}
	g.ResetStyle()
}
