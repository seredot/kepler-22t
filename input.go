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
			} else if key == tcell.KeyLeft || keyRune == 'a' {
				g.player.direction(-1, 0)
			} else if key == tcell.KeyRight || keyRune == 'd' {
				g.player.direction(1, 0)
			} else if key == tcell.KeyUp || keyRune == 'w' {
				g.player.direction(0, -1)
			} else if key == tcell.KeyDown || keyRune == 's' {
				g.player.direction(0, 1)
			} else if keyRune == ' ' {
				g.fire()
			}
		case *tcell.EventMouse:
			x, y := ev.Position()
			button := ev.Buttons()

			// Only process button events, not wheel events
			button &= tcell.ButtonMask(0xff)

			g.mouseX = x
			g.mouseY = y

			switch ev.Buttons() {
			case tcell.ButtonPrimary:
				g.fire()
			}
		}
	}

	return true
}
