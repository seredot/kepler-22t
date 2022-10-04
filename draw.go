package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/seredot/trash/style"
	"gonum.org/v1/gonum/interp"
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

func (g *Game) drawTextTransparent(x, y int, text string) {
	for i, r := range []rune(text) {
		g.PatchChar(x+i, y, r)
	}
}

func (g *Game) drawText(x, y int, text string) {
	for i, r := range []rune(text) {
		g.PutChar(x+i, y, r)
	}
}

func (g *Game) drawHud() {
	g.ResetStyle()

	w := float64(g.width - 1)
	rs := []float64{102, 223, 162, 75, 55}
	gs := []float64{44, 100, 86, 43, 30}
	bs := []float64{55, 100, 206, 137, 97}
	xs := []float64{0, w * 0.2, w * 0.5, w * 0.8, w}
	pr := interp.PiecewiseLinear{}
	pg := interp.PiecewiseLinear{}
	pb := interp.PiecewiseLinear{}
	pr.Fit(xs, rs)
	pg.Fit(xs, gs)
	pb.Fit(xs, bs)

	for x := 0; x < g.width; x++ {
		g.Background(style.ColorRGB(
			uint64(pr.Predict(float64(x))),
			uint64(pg.Predict(float64(x))),
			uint64(pb.Predict(float64(x)))))
		g.PutChar(x, 0, ' ')
	}

	// Title
	g.drawTextTransparent(2, 0, " Kepler 22t ")
	// Stats
	g.drawTextTransparent(2, g.height-1, fmt.Sprintf(" Fr %d | FPS %0.2f ", g.frame, 1000.0/float64(g.deltaT)))
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
