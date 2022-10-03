package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/seredot/trash/style"
)

func (g *Game) ResetStyle() {
	g.style = g.defStyle
}

func (g *Game) Background(c style.Color) {
	g.style = style.Style(tcell.Style(g.style).Background(tcell.Color(c)))
}

func (g *Game) Foreground(c style.Color) {
	g.style = style.Style(tcell.Style(g.style).Foreground(tcell.Color(c)))
}

func (g *Game) PutChar(x, y int, r rune) {
	g.screen.SetContent(x, y, r, nil, tcell.Style(g.style))
}

func (g *Game) PatchChar(x, y int, r rune) {
	_, _, bgStyle, _ := g.screen.GetContent(x, y)
	fgColor, _, _ := tcell.Style(g.style).Decompose()
	mergedStyle := bgStyle.Foreground(fgColor)
	g.screen.SetContent(x, y, r, nil, tcell.Style(mergedStyle))
}

func (g *Game) clear() {
	// Since all the screen is redrawn every frame,
	// it is not necessary to clear the screen.
	g.screen.SetStyle(tcell.Style(g.defStyle))
}

func (g *Game) sync() {
	g.screen.Show()
}

func (g *Game) screenSize() (width, height int) {
	return g.screen.Size()
}

func (g *Game) drawText(x, y int, text string) {
	for i, r := range []rune(text) {
		g.PutChar(x+i, y, r)
	}
}

func (g *Game) drawHud() {
	// Title
	g.drawText(2, 0, " Tr@sh ")
	// Stats
	g.drawText(2, g.height-1, fmt.Sprintf(" Fr %d | FPS %0.2f ", g.frame, 1000.0/float64(g.deltaT)))
	// Debug log
	//g.drawText(22, g.height-1, fmt.Sprintf(" Color %d %x ", g.screen.Colors(), style.Hsl2Rgb(242, 26, 43)))
}

func (g *Game) isInScreen(x, y int) bool {
	if x >= g.left && x <= g.right && y >= g.top && y <= g.bottom {
		return true
	}

	return false
}

func (g *Game) drawAimPointer() {
	if !g.isInScreen(g.mouseX, g.mouseY) {
		return
	}

	g.ResetStyle()
	g.Foreground(style.ColorPointer)
	g.PatchChar(g.mouseX-1, g.mouseY, '❯')
	g.PatchChar(g.mouseX+1, g.mouseY, '❮')
	g.ResetStyle()
}
