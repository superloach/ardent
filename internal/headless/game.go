package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

// Game is a headless implementation
// of engine.Game.
type Game struct {
	tickFunc func()
	c        *component

	Input
}

// NewGame returns an instantiated game.
func NewGame(
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) *Game {
	return &Game{
		tickFunc: tickFunc,
		c:        new(component),
	}
}

// Run starts up the engine and begins
// running the game.
func (g *Game) Run() error {
	for {
		g.tickFunc()
	}
	return nil
}

// Component returns an ebiten component factory.
func (g Game) Component() engine.Component {
	return g.c
}

// AddRenderer adds a renderer to the draw stack.
func (g *Game) AddRenderer(renderer ...engine.Renderer) {
	// NOOP
}

func (g *Game) AddIsoRenderer(isoRenderer ...engine.IsoRenderer) {

}

// IsFullscreen returns the fullscreen state of the game.
func (g Game) IsFullscreen() bool {
	return false
}

// SetFullscreen sets the fullscreen state of the game.
func (g Game) SetFullscreen(v bool) {
	// NOOP
}

// IsVsync returns the vsync state of the game.
func (g Game) IsVsync() bool {
	return false
}

// SetVsync sets the vsync state of the game.
func (g Game) SetVsync(v bool) {
	// NOOP
}

// IsFocused returns the focused state of the game.
func (g Game) IsFocused() bool {
	return false
}
