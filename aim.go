package main

import "github.com/seredot/kepler-22t/style"

func (g *Game) drawAimPointer() {
	if g.OutOfScreen(g.mouseX, g.mouseY) {
		return
	}

	g.ResetStyle()
	g.Foreground(style.ColorPointer)
	g.PatchChar(g.mouseX-1, g.mouseY, '❯')
	g.PatchChar(g.mouseX+1, g.mouseY, '❮')
	g.ResetStyle()
}
