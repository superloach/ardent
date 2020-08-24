package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	renderer engine.Renderer
	image    engine.Image
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {
			image.Translate(1, 0)
			renderer.AddImage(image)
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
	image, _ = component.NewImageFromPath("ebin.jpeg")

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
