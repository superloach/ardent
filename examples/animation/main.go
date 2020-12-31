package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	counter, state uint
	animation      engine.Animation
	animations     = []string{"sw", "se", "nw", "ne"}
)

func main() {
	const w = 57 * 4
	const h = 65 * 4

	// create new game instance
	game := ardent.NewGame(
		"Animation",
		w, h, // 854, 480,
		engine.FlagResizable,
		// tick function
		func() {
			// change animation every 120 ticks
			if counter%120 == 0 && animation != nil {
				animation.SetState(animations[state%4])
				state++
			}
			counter++
		},
		// layout handler
		engine.LayoutFill,
	)

	// create new renderer and animation
	renderer := game.NewRenderer()

	assetutil.CreateAssets("./examples/animation")

	animation, err := game.NewAnimationFromAssetPath("./examples/animation/animation.asset")
	if err != nil {
		log.Fatal(err)
	}

	animation.SetState(animations[0])
	animation.Scale(4, 4)

	// add animation to renderer
	renderer.AddImage(animation)

	// add renderer to game and start game
	game.AddRenderer(renderer)

	err = game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
