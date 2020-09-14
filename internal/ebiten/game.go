package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
)

// Game is an ebiten implementation
// of engine.Game.
type Game struct {
	tickFunc   func()
	layoutFunc func(int, int) (int, int)

	c         *component
	renderers []engine.Renderer

	Input
}

// NewGame returns an instantiated game.
func NewGame(
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) *Game {
	return &Game{
		tickFunc:   tickFunc,
		layoutFunc: layoutFunc,
		c:          new(component),
	}
}

// Run starts up the engine and begins
// running the game.
func (g *Game) Run() error {
	ebiten.SetWindowSize(100, 100)
	ebiten.SetWindowTitle("Ebiten")
	ebiten.SetWindowResizable(true)
	ebiten.SetRunnableInBackground(true)

	return ebiten.RunGame(g)
}

// Component returns an ebiten component factory.
func (g Game) Component() engine.Component {
	return g.c
}

// AddRenderer adds a renderer to the draw stack.
func (g *Game) AddRenderer(renderer ...engine.Renderer) {
	g.renderers = append(g.renderers, renderer...)
}

// Layout is called when the window resizes.
func (g Game) Layout(ow, oh int) (int, int) {
	return g.layoutFunc(ow, oh)
}

// Update runs the tick and draw functions.
func (g *Game) Update(screen *ebiten.Image) error {
	g.tickFunc()

	for _, renderer := range g.renderers {
		renderer.(*Renderer).tick()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for _, renderer := range g.renderers {
		renderer.(*Renderer).draw(screen)
	}

	return nil
}

// IsFullscreen returns the fullscreen state of the game.
func (g Game) IsFullscreen() bool {
	return ebiten.IsFullscreen()
}

// SetFullscreen sets the fullscreen state of the game.
func (g Game) SetFullscreen(v bool) {
	ebiten.SetFullscreen(v)
}

// IsVsync returns the vsync state of the game.
func (g Game) IsVsync() bool {
	return ebiten.IsVsyncEnabled()
}

// SetVsync sets the vsync state of the game.
func (g Game) SetVsync(v bool) {
	ebiten.SetVsyncEnabled(v)
}

// IsFocused returns the focused state of the game.
func (g Game) IsFocused() bool {
	return ebiten.IsFocused()
}
