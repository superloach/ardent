package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	renderer engine.Renderer
	image    engine.Image

	x float64
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {
			x++
			image.Translate(x, 0)
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

	// add image to renderer
	renderer.AddImage(image)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
