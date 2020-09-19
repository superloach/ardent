package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	game    engine.Game
	camera  engine.Camera
	stripes engine.Image
	x, y    float64
)

// tick function
func tick() {
	if game.IsKeyPressed(engine.KeyW) {
		y -= 2
	} else if game.IsKeyPressed(engine.KeyS) {
		y += 2
	}

	if game.IsKeyPressed(engine.KeyA) {
		x -= 2
	} else if game.IsKeyPressed(engine.KeyD) {
		x += 2
	}

	stripes.Translate(x, y)
	camera.LookAt(x, y)
}

func main() {
	// create new game instance
	game = ardent.NewGame(
		"Camera",
		854,
		480,
		engine.FlagResizable,
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

	// create new camera
	camera = component.NewCamera()

	// create new atlas from asset file
	atlas, _ := component.NewAtlasFromAssetPath("../atlas/atlas.asset")

	// get atlas subimages
	stripes = atlas.GetImage("stripes")
	swirls := atlas.GetImage("swirls")
	blocks := atlas.GetImage("blocks")

	// set image positions
	swirls.Translate(128, 0)
	blocks.Translate(128, 128)

	// add images to renderer
	renderer.AddImage(swirls, blocks, stripes)

	// add camera to renderer
	renderer.SetCamera(camera)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
