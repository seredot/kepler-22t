package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen   tcell.Screen
	defStyle tcell.Style
	width    int
	height   int
}

func (g *Game) calcScreenSize() {
	screenWidth, screenHeight := g.screen.Size()
	w, h := 80, 40

	if w > screenWidth {
		w = screenWidth
	}

	if h > screenHeight {
		h = screenHeight
	}

	g.width = w
	g.height = h
}

func (g *Game) drawBorders() {
	for x := 1; x < g.width-1; x++ {
		g.screen.SetContent(x, 0, tcell.RuneHLine, nil, g.defStyle)
		g.screen.SetContent(x, g.height-1, tcell.RuneHLine, nil, g.defStyle)
	}

	for y := 1; y < g.height-1; y++ {
		g.screen.SetContent(0, y, tcell.RuneVLine, nil, g.defStyle)
		g.screen.SetContent(g.width-1, y, tcell.RuneVLine, nil, g.defStyle)
	}

	g.screen.SetContent(0, 0, tcell.RuneULCorner, nil, g.defStyle)
	g.screen.SetContent(0, g.height-1, tcell.RuneLLCorner, nil, g.defStyle)
	g.screen.SetContent(g.width-1, 0, tcell.RuneURCorner, nil, g.defStyle)
	g.screen.SetContent(g.width-1, g.height-1, tcell.RuneLRCorner, nil, g.defStyle)
}

func (g *Game) gameLoop() {
	g.defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset).Blink(true)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	g.screen = s
	s.SetStyle(g.defStyle)
	s.Clear()

	// Character coords
	x, y := 1, 1

	for {
		ev := s.PollEvent()
		g.calcScreenSize()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			key := ev.Key()

			if key == tcell.KeyEscape || key == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			} else if key == tcell.KeyLeft {
				x -= 1
			} else if key == tcell.KeyRight {
				x += 1
			} else if key == tcell.KeyUp {
				y -= 1
			} else if key == tcell.KeyDown {
				y += 1
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		}

		if x < 1 {
			x = 1
		}
		if x > g.width-2 {
			x = g.width - 2
		}
		if y < 1 {
			y = 1
		}
		if y > g.height-2 {
			y = g.height - 2
		}

		s.Clear()
		g.drawBorders()
		s.SetContent(x, y, tcell.RuneDiamond, nil, g.defStyle)
		s.Sync()
	}
}

func main() {
	g := Game{}
	g.gameLoop()
}
