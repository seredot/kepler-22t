package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	opensimplex "github.com/ojrac/opensimplex-go"
	"github.com/seredot/trash/style"
)

type Game struct {
	screen   tcell.Screen
	defStyle style.Style
	style    style.Style
	width    int // screen width
	height   int // screen height
	left     int // left most playable area
	right    int // right most playable area
	top      int // top most playable area
	bottom   int //  bottom most playable area
	mouseX   int
	mouseY   int

	frame  int
	deltaT int64
	noise  opensimplex.Noise

	player  *Player
	enemies []*Enemy
}

func (g *Game) init() {
	g.defStyle = style.DefaultStyle()

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	g.screen = s
	s.EnableMouse()
	g.calcScreenSize()
	g.clear()

	// Player initials
	g.player = NewPlayer(g, 10, 10)
	g.enemies = []*Enemy{}

	// Random noise generator
	g.noise = opensimplex.NewNormalized(110783)
}

func (g *Game) calcScreenSize() {
	screenWidth, screenHeight := g.screenSize()
	w, h := 80, 40

	if w > screenWidth {
		w = screenWidth
	}

	if h > screenHeight {
		h = screenHeight
	}

	g.width = w
	g.height = h

	g.left = 1
	g.right = g.width - 2
	g.top = 1
	g.bottom = g.height - 2
}

func (g *Game) spawnEnemy() {
	if g.frame%100 == 0 {
		g.enemies = append(g.enemies, NewEnemy(g))
	}
}

func (g *Game) drawEnemies() {
	for _, e := range g.enemies {
		e.draw()
	}
}

func (g *Game) loop() {
	lastFrameT := time.Now().UnixMilli()
	g.deltaT = 10

	for {
		if !g.processInput() {
			return
		}

		g.spawnEnemy()
		g.calcScreenSize()
		g.clear()
		g.drawBorders()
		g.drawTerrain()
		g.drawEnemies()
		g.player.draw()
		g.drawAimPointer()
		g.drawHud()

		g.frame++
		g.sync()

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
