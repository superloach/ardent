package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	counter, state uint
	animation      engine.Animation
	animations     = []string{"sw", "se", "nw", "ne"}
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Animation",
		854,
		480,
		engine.FlagResizable,
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {
			// change animation every 120 ticks
			if counter%120 == 0 {
				animation.SetState(animations[state%4])
				state++
			}
			counter++
		},
		// layout function
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	// get component factory
	component := game.Component()

	// create new renderer and animation
	renderer := component.NewRenderer()
	animation, _ = component.NewAnimationFromAssetPath("animation.asset")
	animation.Scale(4, 4)

	// add animation to renderer
	renderer.AddImage(animation)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
