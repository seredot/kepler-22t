package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/seredot/kepler-22t/color"
)

type Cell struct {
	fgColor color.Color
	bgColor color.Color
	char    rune
}

func ColorConv(c color.Color) tcell.Color {
	return tcell.Color(uint64(tcell.ColorIsRGB) | uint64(tcell.ColorValid) | uint64(c.R*255.0)<<16 | uint64(c.G*255.0)<<8 | uint64(c.B*255.0))
}

func (g *Game) sync() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			cell := g.getCell(x, y)
			g.screen.SetContent(x, y, cell.char, nil,
				tcell.Style(tcell.StyleDefault).Background(ColorConv(cell.bgColor)).Foreground(ColorConv(cell.fgColor)))
		}
	}

	g.screen.Show()
}

func (g *Game) screenSize() (width, height int) {
	return g.screen.Size()
}

func (g *Game) ResetStyle() {
	g.bgColor = color.ColorTransparent
	g.fgColor = color.ColorWhite
}

func (g *Game) Background(c color.Color) {
	g.bgColor = c
}

func (g *Game) Foreground(c color.Color) {
	g.fgColor = c
}

func (g *Game) OutOfScreen(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return true
	}

	return false
}

func (g *Game) getCell(x, y int) *Cell {
	return &(g.cells[y*g.width+x])
}

func (g *Game) PutChar(x, y int, r rune) {
	if g.OutOfScreen(x, y) {
		return
	}

	cell := &g.cells[y*g.width+x]
	cell.char = r
	cell.bgColor = cell.bgColor.Blend(g.bgColor)
	cell.fgColor = cell.fgColor.Blend(g.fgColor)
}

func (g *Game) PatchChar(x, y int, r rune) {
	if g.OutOfScreen(x, y) {
		return
	}

	cell := &g.cells[y*g.width+x]
	cell.char = r
	cell.bgColor = cell.bgColor.Blend(g.bgColor)
	cell.fgColor = cell.fgColor.Blend(g.fgColor)
}

func (g *Game) PutColor(x, y int) {
	if g.OutOfScreen(x, y) {
		return
	}

	cell := &g.cells[y*g.width+x]
	cell.bgColor = cell.bgColor.Blend(g.bgColor)
	cell.fgColor = cell.fgColor.Blend(g.fgColor)
}

func (g *Game) DrawText(x, y int, text string) {
	for i, r := range []rune(text) {
		g.PatchChar(x+i, y, r)
	}
}
