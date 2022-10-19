package main

import (
	"fmt"
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

	// Title
	g.DrawTextTransparent(1, 0, "Kepler 22t")
	// Score
	score := fmt.Sprintf("☠%d", g.score)
	g.DrawTextTransparent(g.width-1-utf8.RuneCountInString(score), 0, score)
	// Stats
	g.DrawTextTransparent(2, g.height-1, fmt.Sprintf(" Fr %d | FPS %0.2f ", g.frame, float64(time.Second/g.deltaT)))
	// Debug log
	//g.drawText(22, g.height-1, fmt.Sprintf(" Color %d %x ", g.screen.Colors(), style.Hsl2Rgb(242, 26, 43)))
}
