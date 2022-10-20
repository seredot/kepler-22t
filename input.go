package main

import (
	"github.com/gdamore/tcell/v2"
)

func (g *Game) processInput() bool {
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
				return false
			}

			if g.state == Playing {
				if key == tcell.KeyLeft || keyRune == 'a' {
					g.player.direction(-1, 0)
				} else if key == tcell.KeyRight || keyRune == 'd' {
					g.player.direction(1, 0)
				} else if key == tcell.KeyUp || keyRune == 'w' {
					g.player.direction(0, -1)
				} else if key == tcell.KeyDown || keyRune == 's' {
					g.player.direction(0, 1)
				}
			}

			if g.state == GameOver {
				if key == tcell.KeyEnter {
					g.reset()
				}
			}
		case *tcell.EventMouse:
			x, y := ev.Position()
			button := ev.Buttons()

			g.mouseX = x
			g.mouseY = y
			g.mouseDown = button&tcell.Button1 != 0
		}
	}

	return true
}
