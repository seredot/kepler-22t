package main

import (
	"log"
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
	opensimplex "github.com/ojrac/opensimplex-go"
	"github.com/seredot/kepler-22t/style"
)

type Game struct {
	// Canvas
	canvas   Canvas
	screen   tcell.Screen
	defStyle style.Style
	style    style.Style

	coords Coords
	width  int // screen width
	height int // screen height
	left   int // left most playable area
	right  int // right most playable area
	top    int // top most playable area
	bottom int //  bottom most playable area
	mouseX int
	mouseY int

	// Timing
	timing Timing
	frame  int
	deltaT int64

	// Objects
	player  *Player
	enemies []*Enemy
	bullets []*Bullet

	// Misc
	noise opensimplex.Noise
}

func NewGame() *Game {
	g := &Game{}

	g.timing = g
	g.coords = g
	g.canvas = g

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
	mag := math.Sqrt(dx*dx + dy*dy)
	dx /= mag
	dy /= mag
	g.bullets = append(g.bullets, NewBullet(g.player.x, g.player.y, dx, dy, 30))
}

func (g *Game) spawnEnemy() {
	if g.frame%100 == 0 {
		g.enemies = append(g.enemies, NewEnemy(g))
	}
}

func (g *Game) moveEnemies() {
	for _, e := range g.enemies {
		e.move(g.timing)
	}
}

func (g *Game) drawEnemies() {
	for _, e := range g.enemies {
		e.draw(g.canvas)
	}
}

func (g *Game) moveBullets() {
	nextBullets := make([]*Bullet, 0, len(g.bullets))

	for _, b := range g.bullets {
		b.move(g.timing, g.coords)
		if !b.removed {
			nextBullets = append(nextBullets, b)
		}
	}

	g.bullets = nextBullets
}

func (g *Game) drawBullets() {
	for _, b := range g.bullets {
		b.draw(g.canvas)
	}
}

func (g *Game) Loop() {
	lastFrameT := time.Now().UnixMilli()
	g.deltaT = 10

	for {
		if !g.processInput() {
			return
		}

		g.spawnEnemy()
		g.calcScreenSize()
		g.clear()
		g.drawTerrain()
		g.moveEnemies()
		g.drawEnemies()
		g.player.move(g.timing, g.coords)
		g.player.draw(g.canvas)
		g.moveBullets()
		g.drawBullets()
		hitBullets(g.bullets, g.enemies)
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
