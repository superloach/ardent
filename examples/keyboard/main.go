package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	game      engine.Game
	animation engine.Animation
	x         float64
)

// tick function
func tick() {
	if game.IsKeyPressed(engine.KeyA) {
		// walk left
		animation.SetState("sw")
		x--
	} else if game.IsKeyPressed(engine.KeyD) {
		// walk right
		animation.SetState("se")
		x++
	}

	animation.Translate(x, 0)
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

	// create new renderer and animation
	renderer := component.NewRenderer()
	animation, _ = component.NewAnimationFromAssetPath("../animation/animation.asset")
	animation.Scale(4, 4)
	animation.SetState("sw")

	// add animation to renderer
	renderer.AddImage(animation)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
