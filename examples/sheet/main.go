package main

import (
	"github.com/split-cube-studios/ardent"
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {
		},
		// layout function
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	// get component factory
	component := game.Component()

	// create new renderer and image
	renderer = component.NewRenderer()

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
