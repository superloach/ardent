package main

import (
	"github.com/split-cube-studios/ardent"
)

func main() {
	game := ardent.NewGame(
		ardent.EBITEN,
		func() {
			// TODO
		},
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	component := game.Component()
	renderer := component.NewRenderer()

	game.AddRenderer(renderer)
	game.Run()
}
