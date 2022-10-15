package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/seredot/kepler-22t/color"
)

type Canvas interface {
	Coords

	ResetStyle()
	Background(c color.Color)
	Foreground(c color.Color)
	OutOfScreen(x, y int) bool
	PutChar(x, y int, r rune)
	PatchChar(x, y int, r rune)
	DrawTextTransparent(x, y int, text string)
	DrawText(x, y int, text string)
}

var DefaultStyle = tcell.Style(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite).Blink(false))

func ColorConv(c color.Color) tcell.Color {
	return tcell.Color(uint64(tcell.ColorIsRGB) | uint64(tcell.ColorValid) | uint64(c.R*255.0)<<16 | uint64(c.G*255.0)<<8 | uint64(c.B*255.0))
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

func (g *Game) ResetStyle() {
	g.style = g.defStyle
}

func (g *Game) Background(c color.Color) {
	g.style = g.style.Background(ColorConv(c))
}

func (g *Game) Foreground(c color.Color) {
	g.style = g.style.Foreground(ColorConv(c))
}

func (g *Game) OutOfScreen(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return true
	}

	return false
}

func (g *Game) PutChar(x, y int, r rune) {
	if g.OutOfScreen(x, y) {
		return
	}

	g.screen.SetContent(x, y, r, nil, tcell.Style(g.style))
}

func (g *Game) PatchChar(x, y int, r rune) {
	if g.OutOfScreen(x, y) {
		return
	}

	_, _, bgStyle, _ := g.screen.GetContent(x, y)
	fgColor, _, _ := tcell.Style(g.style).Decompose()
	mergedStyle := bgStyle.Foreground(fgColor)
	g.screen.SetContent(x, y, r, nil, tcell.Style(mergedStyle))
}

func (g *Game) DrawTextTransparent(x, y int, text string) {
	for i, r := range []rune(text) {
		g.PatchChar(x+i, y, r)
	}
}

func (g *Game) DrawText(x, y int, text string) {
	for i, r := range []rune(text) {
		g.PutChar(x+i, y, r)
	}
}
