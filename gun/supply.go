package gun

import (
	"github.com/seredot/kepler-22t/color"
	"github.com/seredot/kepler-22t/object"
	"github.com/seredot/kepler-22t/screen"
)

type SupplyType int

const (
	Health SupplyType = iota
	SemiAutomatic
	MachineGun
	GatlingGun
	RailGun
	FlameThrower
	PlasmaGun
	Nuke
	Freezer
	TripleDamage

	NoSupply
)

type SupplyBox struct {
	object.Object
	Type SupplyType
	Gun  Gun
}

var (
	HealthBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorWhite,
			FgColor: color.ColorCrossRed,
			Sprite:  '✚',
		},
		Type: Health,
	}
	SemiAutomaticBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBullet,
			Sprite:  'S',
		},
		Type: Health,
		Gun:  NewMachineGun(),
	}
	MachineGunBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBullet,
			Sprite:  'M',
		},
		Type: Health,
		Gun:  NewMachineGun(),
	}
	GatlingGunBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBullet,
			Sprite:  'G',
		},
		Type: Health,
		Gun:  NewMachineGun(),
	}
	RailGunBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBullet,
			Sprite:  'R',
		},
		Type: Health,
		Gun:  NewMachineGun(),
	}
	FlameThrowerBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBullet,
			Sprite:  '⌬',
		},
		Type: Health,
		Gun:  NewMachineGun(),
	}
	PlasmaGunBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBullet,
			Sprite:  '⚛',
		},
		Type: Health,
		Gun:  NewMachineGun(),
	}
	NukeBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBlack,
			Sprite:  '☢',
		},
		Type: Health,
	}
	FreezerBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBlack,
			Sprite:  '❄',
		},
		Type: Health,
	}
	TripleDamageBox = SupplyBox{
		Object: object.Object{
			BgColor: color.ColorBox,
			FgColor: color.ColorBlack,
			Sprite:  '3',
		},
		Type: Health,
	}
)

func (s SupplyBox) Draw(c screen.Canvas) {
	c.ResetStyle()
	c.Background(s.BgColor)
	c.PutColor(s.ScrX()-1, s.ScrY())
	c.PutColor(s.ScrX()+1, s.ScrY())
	c.Foreground(s.FgColor)
	c.PatchChar(s.ScrX(), s.ScrY(), s.Sprite)
	c.ResetStyle()
}
