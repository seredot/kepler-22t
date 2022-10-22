package game

import "github.com/seredot/kepler-22t/color"

func (g *Game) drawAimPointer() {
	if g.OutOfScreen(g.mouseX, g.mouseY) {
		return
	}

	g.ResetStyle()
	g.Foreground(color.ColorPointer)
	g.PatchChar(g.mouseX-1, g.mouseY, '❯')
	g.PatchChar(g.mouseX+1, g.mouseY, '❮')
	g.ResetStyle()
}
