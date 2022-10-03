package main

import (
	"github.com/gdamore/tcell/v2"
)

func (g *Game) processKeyboard() bool {
	for g.screen.HasPendingEvent() {
		ev := g.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			key := ev.Key()
			keyRune := ev.Rune()

			if key == tcell.KeyEscape || key == tcell.KeyCtrlC {
				g.screen.Clear()
				g.screen.Fini()
				return false
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

	return true
}
