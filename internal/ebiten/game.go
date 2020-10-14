package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
)

// Game is an ebiten implementation
// of engine.Game.
type Game struct {
	title string
	w, h  int
	flags byte

	tickFunc   func()
	layoutFunc func(int, int) (int, int)

	renderers []engine.Renderer

	component
	Input
}

// NewGame returns an instantiated game.
func NewGame(
	title string,
	w, h int,
	flags byte,
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) *Game {
	return &Game{
		title:      title,
		w:          w,
		h:          h,
		flags:      flags,
		tickFunc:   tickFunc,
		layoutFunc: layoutFunc,
	}
}

// Run starts up the engine and begins
// running the game.
func (g *Game) Run() error {
	ebiten.SetWindowSize(g.w, g.h)
	ebiten.SetWindowTitle(g.title)
	ebiten.SetWindowResizable(g.flags&engine.FlagResizable > 0)
	ebiten.SetRunnableInBackground(g.flags&engine.FlagRunsInBackground > 0)

	return ebiten.RunGame(g)
}

// AddRenderer adds a renderer to the draw stack.
func (g *Game) AddRenderer(renderer ...engine.Renderer) {
	g.renderers = append(g.renderers, renderer...)
}

// Layout is called when the window resizes.
func (g *Game) Layout(ow, oh int) (int, int) {
	g.w, g.h = g.layoutFunc(ow, oh)
	return g.w, g.h
}

// Update runs the tick and draw functions.
func (g *Game) Update(screen *ebiten.Image) error {
	g.tickFunc()

	for _, renderer := range g.renderers {
		renderer.Tick()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for _, renderer := range g.renderers {
		renderer.SetViewport(g.w, g.h)
		switch renderer.(type) {
		case *Renderer:
			renderer.(*Renderer).draw(screen)
		case *IsoRenderer:
			renderer.(*IsoRenderer).draw(screen)
		}
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
