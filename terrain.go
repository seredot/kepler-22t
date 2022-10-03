package main

import "github.com/seredot/trash/style"

func (g *Game) drawTerrain() {
	g.ResetStyle()
	for x := g.left; x <= g.right; x++ {
		for y := g.top; y <= g.bottom; y++ {
			g.Background(style.Color(style.Hsl2Rgb(242, 26, 43)))
			g.Foreground(style.Color(style.Hsl2Rgb(242, 26, 43)))
			g.PutChar(x, y, ' ')
		}
	}
	g.ResetStyle()
}
