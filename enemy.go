package main

import "github.com/seredot/kepler-22t/style"

type Enemy struct {
	Object
	energy float64
}

func NewEnemy(game *Game) *Enemy {
	x := float64(game.left) + float64(game.right-game.left)*game.noise.Eval2(0, float64(game.frame))
	y := float64(game.top) + float64(game.bottom-game.top)*game.noise.Eval2(float64(game.frame), 0)

	e := &Enemy{
		Object: Object{
			x:      x,
			y:      y,
			sprite: '✹',
			color:  style.ColorEnemy,
		},
		energy: 100.0,
	}

	return e
}
