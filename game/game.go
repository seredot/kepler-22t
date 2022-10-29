package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	opensimplex "github.com/ojrac/opensimplex-go"
	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/gun"
	"github.com/seredot/kepler-22t/object"
	"github.com/seredot/kepler-22t/screen"
)

type GameState int

const (
	Playing GameState = iota
	GameOver
)

const MaxFPS = 50.0

type Game struct {
	// Canvas
	canvas  screen.Canvas
	cells   []Cell
	screen  tcell.Screen
	fgColor color.Color
	bgColor color.Color

	coords    screen.Coords
	width     int // screen width
	height    int // screen height
	left      int // left most playable area
	right     int // right most playable area
	top       int // top most playable area
	bottom    int // bottom most playable area
	mouseX    int
	mouseY    int
	mouseDown bool

	// Timing
	frame       int           // frame counter
	totalT      time.Duration // total play
	simuDelta   time.Duration // since simulation iter
	renderDelta time.Duration // since frame render

	// Objects
	player   *object.Player
	aliens   []*object.Alien
	supplies []*gun.SupplyBox
	bullets  []*object.Bullet
	effects  []*object.Effect

	// Misc
	state  GameState
	score  int
	health float64
	ammo   int
	gun    gun.Gun
	fireT  time.Time
	noise  opensimplex.Noise
}

func NewGame() *Game {
	g := &Game{}

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

	// Random noise generator
	g.noise = opensimplex.NewNormalized(110783)

	// Game initials
	g.reset()

	return g
}

func (g *Game) reset() {
	g.health = 100.0
	g.score = 0
	g.ammo = 30
	g.gun = gun.NewMachineGun()
	g.mouseDown = false
	g.player = object.NewPlayer(10, 10)
	g.aliens = []*object.Alien{}
	g.supplies = []*gun.SupplyBox{}
	g.bullets = []*object.Bullet{}
	g.state = Playing
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

func (g *Game) Frame() int {
	return g.frame
}

func (g *Game) DeltaT() time.Duration {
	return g.simuDelta
}

func (g *Game) Noise() opensimplex.Noise {
	return g.noise
}

func (g *Game) handleTrigger() {
	if !g.mouseDown {
		return
	}

	if time.Since(g.fireT) < g.gun.Delay() {
		return
	}

	if g.ammo != 0 {
		g.fireT = time.Now()
		g.ammo--

		g.bullets = append(g.bullets, g.gun.Fire(g)...)
		g.addEffects(object.NewGunFlash(g.player.X, g.player.Y)...)
	}

	if g.ammo <= 0 {
		g.gun = gun.NewPistol()
		g.ammo = -1
	}
}

func (g *Game) spawnAlien() {
	if g.frame%300 == 0 {
		g.aliens = append(g.aliens, object.NewAlien(g))
	}
}

func (g *Game) spawnSupply() {
	if g.frame%500 != 0 {
		return
	}

	x := float64(rand.Int() % g.Width())
	y := float64(rand.Int() % g.Height())

	boxes := []gun.SupplyBox{
		gun.HealthBox,
		gun.SemiAutomaticBox,
		gun.MachineGunBox,
		gun.GatlingGunBox,
		gun.RailGunBox,
		gun.FlameThrowerBox,
		gun.PlasmaGunBox,
		gun.NukeBox,
		gun.FreezerBox,
		gun.TripleDamageBox,
	}

	box := boxes[rand.Int()%len(boxes)]
	box.X = x
	box.Y = y

	g.supplies = append(g.supplies, &box)
}

func (g *Game) moveAliens() {
	next := make([]*object.Alien, 0, len(g.aliens))

	for _, a := range g.aliens {
		a.Move(g.coords)
		if !a.Removed {
			next = append(next, a)
		}
	}

	g.aliens = next
}

func (g *Game) drawAliens() {
	for _, a := range g.aliens {
		a.Draw(g.canvas)
	}
}

func (g *Game) drawSupplies() {
	for _, s := range g.supplies {
		s.Draw(g.canvas)
	}
}

func (g *Game) moveBullets() {
	next := make([]*object.Bullet, 0, len(g.bullets))

	for _, b := range g.bullets {
		b.Move(g.canvas)
		if !b.Removed {
			next = append(next, b)
		}
	}

	g.bullets = next
}

func (g *Game) drawBullets() {
	for _, b := range g.bullets {
		b.Draw(g.canvas)
	}
}

func (g *Game) addEffects(elems ...*object.Effect) {
	g.effects = append(g.effects, elems...)
}

func (g *Game) moveEffects() {
	next := make([]*object.Effect, 0, len(g.effects))

	for _, e := range g.effects {
		e.Move(g.canvas)
		if !e.Removed {
			next = append(next, e)
		}
	}

	g.effects = next
}

func (g *Game) drawEffects() {
	for _, e := range g.effects {
		e.Draw(g.canvas)
	}
}

func (g *Game) resetScreen() {
	g.calcScreenSize()
	g.cells = make([]Cell, g.width*g.height)
}

func (g *Game) Loop() {
	lastFrameT := time.Now()
	lastRenderT := time.Now()
	g.simuDelta = 10 * time.Millisecond
	g.renderDelta = 10 * time.Millisecond

	for {
		if !g.processInput() {
			return
		}

		// Simulate
		g.spawnAlien()
		g.spawnSupply()
		g.handleTrigger()
		g.moveAliens()
		g.moveBullets()
		g.player.Move(g.canvas)
		g.checkCollisions()
		g.moveEffects()

		// Render
		if time.Since(lastRenderT) >= (time.Second / MaxFPS) {
			g.renderDelta = time.Since(lastRenderT)
			g.resetScreen()
			g.drawFog()
			g.drawSupplies()
			g.drawAliens()
			g.drawBullets()
			g.player.Draw(g.canvas)
			g.drawEffects()
			g.drawAimPointer()
			g.drawHud()
			g.sync()
			lastRenderT = time.Now()
		}

		// Calculate delta time between simulation frames
		g.frame++
		now := time.Now()
		g.simuDelta = now.Sub(lastFrameT)
		if g.simuDelta > 100*time.Millisecond {
			g.simuDelta = 100 * time.Millisecond
		}
		g.totalT = g.totalT + g.simuDelta
		lastFrameT = now

		// 100 simulation frames
		time.Sleep(10 * time.Millisecond)
	}
}
