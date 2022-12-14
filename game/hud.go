package game

import (
	"fmt"
	"math"
	"time"
	"unicode/utf8"

	"github.com/seredot/kepler-22t/color"
	"gonum.org/v1/gonum/interp"
)

func (g *Game) drawHud() {
	g.ResetStyle()

	w := float64(g.width - 1)
	rs := []float64{102, 223, 162, 75, 55}
	gs := []float64{44, 100, 86, 43, 30}
	bs := []float64{55, 100, 206, 137, 97}
	xs := []float64{0, w * 0.2, w * 0.5, w * 0.8, w}
	pr := interp.ClampedCubic{}
	pg := interp.ClampedCubic{}
	pb := interp.ClampedCubic{}
	pr.Fit(xs, rs)
	pg.Fit(xs, gs)
	pb.Fit(xs, bs)

	for x := 0; x < g.width; x++ {
		g.Background(color.NewColorIntRGB(
			uint64(pr.Predict(float64(x))),
			uint64(pg.Predict(float64(x))),
			uint64(pb.Predict(float64(x)))))
		g.PutChar(x, 0, ' ')
	}

	g.ResetStyle()

	// Title
	g.DrawText(1, 0, "Kepler 22t")

	// Right aligned indicators
	var textX, textY int
	textX = g.width
	printRight := func(s string) {
		textX -= utf8.RuneCountInString(s)
		g.DrawText(textX, textY, s)
	}
	printLeft := func(s string) {
		g.DrawText(textX, textY, s)
		textX += utf8.RuneCountInString(s)
	}

	// Score
	g.Foreground(color.ColorWhite)
	printRight(fmt.Sprintf(" %d ", g.score))
	g.Foreground(color.ColorAlien)
	printRight("☠")
	g.ResetStyle()

	// Health
	printRight(fmt.Sprintf(" %d ", int(math.Round(g.health))))
	g.Foreground(color.ColorCrossRed)
	printRight("✚")
	g.ResetStyle()

	// Ammo
	if g.ammo == -1 {
		printRight(" ∞ ")
	} else {
		printRight(fmt.Sprintf(" %d ", g.ammo))
	}
	g.Foreground(color.ColorAmber)
	printRight("⁍")
	g.ResetStyle()

	// Game over
	if g.state == GameOver {
		var m string
		g.Background(color.ColorRedSpill)
		m = "      GAME OVER       "
		textX = (g.width - utf8.RuneCountInString(m)) / 2
		textY = (g.height / 2)
		g.DrawText(textX, textY, m)
		m = " press enter to start "
		textX = (g.width - utf8.RuneCountInString(m)) / 2
		textY++
		g.DrawText(textX, textY, m)
		g.ResetStyle()
	}

	// Gun
	textX = g.width
	textY = g.height - 1
	g.Foreground(color.ColorAmber)
	printRight(fmt.Sprintf(" %s ", g.gun.Name()))
	g.ResetStyle()

	// Debug
	textX = 1
	textY = g.height - 1
	printLeft(fmt.Sprintf("FPS %5.1f | ", float64(time.Second/g.renderDelta)))
	printLeft(fmt.Sprintf("Aliens %d | ", len(g.aliens)))
	printLeft(fmt.Sprintf("Bullets %d ", len(g.bullets)))
}
