package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen   tcell.Screen
	defStyle tcell.Style
	width    int
	height   int
	player   *Player
	frame    int
	deltaT   int64
}

func (g *Game) init() {
	g.defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset).Blink(false)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	g.screen = s
	g.calcScreenSize()
	s.SetStyle(g.defStyle)
	s.Clear()

	// Player initials
	g.player = NewPlayer(g, 10, 10)
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

func (g *Game) drawText(x, y int, text string) {
	for i, r := range []rune(text) {
		g.screen.SetContent(x+i, y, r, nil, g.defStyle)
	}
}

func (g *Game) drawHud() {
	// Title
	g.drawText(2, 0, " Tr@sh ")
	// Stats
	g.drawText(2, g.height-1, fmt.Sprintf(" Fr %d | FPS %0.2f ", g.frame, 1000.0/float64(g.deltaT)))
}

func (g *Game) processKeyboard() {
	for g.screen.HasPendingEvent() {
		ev := g.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			key := ev.Key()
			keyRune := ev.Rune()

			if key == tcell.KeyEscape || key == tcell.KeyCtrlC {
				g.screen.Fini()
				os.Exit(0)
			} else if key == tcell.KeyLeft {
				g.player.direction(-1, 0)
			} else if key == tcell.KeyRight {
				g.player.direction(1, 0)
			} else if key == tcell.KeyUp {
				g.player.direction(0, -1)
			} else if key == tcell.KeyDown {
				g.player.direction(0, 1)
			} else if keyRune == ' ' {
				g.player.fire()
			}
		}
	}
}

func (g *Game) loop() {
	lastFrameT := time.Now().UnixMilli()
	g.deltaT = 10

	for {
		g.processKeyboard()
		g.calcScreenSize()
		g.screen.Clear()
		g.drawBorders()
		g.player.draw()
		g.drawHud()

		g.frame++
		g.screen.Sync()

		// Calculate delta time between frames
		now := time.Now().UnixMilli()
		g.deltaT = now - lastFrameT
		lastFrameT = now
		if g.deltaT > 100 {
			g.deltaT = 100
		}

		// Limit to 30 fps
		time.Sleep(time.Duration(33-g.deltaT) * time.Millisecond)
	}
}
