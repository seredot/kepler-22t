package main

import "github.com/seredot/trash/style"

type Enemy struct {
	Object
	energy float64
}

func NewEnemy(game *Game) *Enemy {
	x := float64(game.left) + float64(game.width)*game.noise.Eval2(0, float64(game.frame))
	y := float64(game.top) + float64(game.height)*game.noise.Eval2(float64(game.frame), 0)

	e := &Enemy{
		Object: Object{
			game:   game,
			x:      x,
			y:      y,
			sprite: '✹',
			color:  style.ColorEnemy,
		},
		energy: 100.0,
	}

	return e
}
