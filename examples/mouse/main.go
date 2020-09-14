package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	game    engine.Game
	stripes engine.Image
)

// tick function
func tick() {
	cx, cy := game.CursorPosition()
	w, h := stripes.Size()
	stripes.Translate(float64(cx-w/2), float64(cy-h/2))
}

func main() {
	// create new game instance
	game = ardent.NewGame(
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		tick,
		// layout function
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	// get component factory
	component := game.Component()

	// create new renderer
	renderer := component.NewRenderer()

	// create new atlas from asset file
	atlas, _ := component.NewAtlasFromAssetPath("../atlas/atlas.asset")

	// get atlas subimage
	stripes = atlas.GetImage("stripes")

	// add image to renderer
	renderer.AddImage(stripes)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
