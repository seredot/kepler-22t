package style

import "github.com/gdamore/tcell/v2"

type Color tcell.Color
type Style tcell.Style

const (
	ColorReset      = Color(tcell.ColorReset)
	ColorPlayer     = Color(tcell.ColorBlue)
	ColorBullet     = Color(tcell.ColorYellow)
	ColorBackground = Color(tcell.ColorBlack)
	ColorForeground = Color(tcell.ColorWhite)
)

func DefaultStyle() Style {
	return Style(tcell.StyleDefault.Background(tcell.Color(ColorBackground)).Foreground(tcell.Color(ColorForeground)).Blink(false))
}
