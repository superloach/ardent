package main

import (
	"math"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Atlas",
		854,
		480,
		engine.FlagResizable,
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {},
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
	atlas, _ := component.NewAtlasFromAssetPath("atlas.asset")

	// get atlas subimages
	stripes := atlas.GetImage("stripes")
	swirls := atlas.GetImage("swirls")
	blocks := atlas.GetImage("blocks")

	// set image positions
	stripes.Rotate(math.Pi / 3)
	swirls.Translate(128, 0)
	blocks.Translate(128, 128)
	blocks.Scale(0.5, 2)

	// add images to renderer
	renderer.AddImage(stripes, swirls, blocks)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
