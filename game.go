package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	opensimplex "github.com/ojrac/opensimplex-go"
	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/vector"
)

type Game struct {
	// Canvas
	canvas  Canvas
	cells   []Cell
	screen  tcell.Screen
	fgColor color.Color
	bgColor color.Color

	coords Coords
	width  int // screen width
	height int // screen height
	left   int // left most playable area
	right  int // right most playable area
	top    int // top most playable area
	bottom int // bottom most playable area
	mouseX int
	mouseY int

	// Timing
	timing Timing
	frame  int
	totalT time.Duration
	deltaT time.Duration

	// Objects
	player  *Player
	aliens  []*Alien
	bullets []*Bullet
	effects []*Effect

	// Misc
	noise opensimplex.Noise
}

func NewGame() *Game {
	g := &Game{}

	g.timing = g
	g.coords = g
	g.canvas = g

	// Initialize screen
	g.cells = []Cell{}

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
	g.ResetStyle()

	// Player initials
	g.player = NewPlayer(g, 10, 10)
	g.aliens = []*Alien{}
	g.bullets = []*Bullet{}

	// Random noise generator
	g.noise = opensimplex.NewNormalized(110783)

	return g
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

	g.left = 0
	g.right = g.width - 1
	g.top = 1
	g.bottom = g.height - 2
}

func (g *Game) fire() {
	dx := float64(g.mouseX) - g.player.x
	dy := float64(g.mouseY) - g.player.y
	dx, dy = vector.Norm(dx, dy)
	g.bullets = append(g.bullets, NewBullet(g.player.x, g.player.y, dx, dy, 30))
	g.addEffects(NewGunFlash(g.player.x, g.player.y)...)
}

func (g *Game) spawnAlien() {
	if g.frame%100 == 0 {
		g.aliens = append(g.aliens, NewAlien(g))
	}
}

func (g *Game) moveAliens() {
	next := make([]*Alien, 0, len(g.aliens))

	for _, a := range g.aliens {
		a.move(g.timing, g.coords)
		if !a.removed {
			next = append(next, a)
		}
	}

	g.aliens = next
}

func (g *Game) drawAliens() {
	for _, a := range g.aliens {
		a.draw(g.canvas)
	}
}

func (g *Game) moveBullets() {
	next := make([]*Bullet, 0, len(g.bullets))

	for _, b := range g.bullets {
		b.move(g.timing, g.coords)
		if !b.removed {
			next = append(next, b)
		}
	}

	g.bullets = next
}

func (g *Game) drawBullets() {
	for _, b := range g.bullets {
		b.draw(g.canvas)
	}
}

func (g *Game) addEffects(elems ...*Effect) {
	g.effects = append(g.effects, elems...)
}

func (g *Game) moveEffects() {
	next := make([]*Effect, 0, len(g.effects))

	for _, e := range g.effects {
		e.move(g.timing, g.coords)
		if !e.removed {
			next = append(next, e)
		}
	}

	g.effects = next
}

func (g *Game) drawEffects() {
	for _, e := range g.effects {
		e.draw(g.canvas)
	}
}

func (g *Game) resetScreen() {
	g.calcScreenSize()
	g.cells = make([]Cell, g.width*g.height)
}

func (g *Game) Loop() {
	lastFrameT := time.Now()
	g.deltaT = 10 * time.Millisecond

	for {
		if !g.processInput() {
			return
		}

		g.resetScreen()
		g.spawnAlien()
		g.moveAliens()
		g.drawAliens()
		g.moveBullets()
		g.drawBullets()
		g.player.move(g.timing, g.coords)
		g.player.draw(g.canvas)
		hitBullets(g, g.bullets, g.aliens)
		g.moveEffects()
		g.drawEffects()
		g.drawFog()
		g.drawAimPointer()
		g.drawHud()

		g.frame++
		g.sync()

		// Calculate delta time between frames
		now := time.Now()
		g.deltaT = now.Sub(lastFrameT)
		g.totalT = g.totalT + g.deltaT
		lastFrameT = now
		if g.deltaT > 100*time.Millisecond {
			g.deltaT = 100 * time.Millisecond
		}

		// Limit to 30 fps
		time.Sleep((33 * time.Millisecond) - g.deltaT)
	}
}
