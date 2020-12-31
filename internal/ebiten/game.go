//+build !headless

package ebiten

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/engine"
)

// Game is an ebiten implementation
// of engine.Game.
type Game struct {
	title  string
	vw, vh int
	w, h   int
	ow, oh int
	flags  byte

	tickFunc   func()
	layoutFunc engine.LayoutFunc

	renderers []engine.Renderer

	*component
	Input
}

// NewGame returns an instantiated game.
func NewGame(
	title string,
	w, h int,
	flags byte,
	tickFunc func(),
	layoutFunc engine.LayoutFunc,
) *Game {
	return &Game{
		title: title,
		vw:    w,
		vh:    h,
		w:     w,
		h:     h,
		flags: flags,

		tickFunc:   tickFunc,
		layoutFunc: layoutFunc,

		component: newComponent(),
	}
}

// Run starts up the engine and begins
// running the game.
func (g *Game) Run() error {
	ebiten.SetWindowSize(g.vw, g.vh)
	ebiten.SetWindowTitle(g.title)
	ebiten.SetWindowResizable(g.flags&engine.FlagResizable > 0)
	ebiten.SetRunnableOnUnfocused(g.flags&engine.FlagRunsInBackground > 0)

	return ebiten.RunGame(g)
}

// AddRenderer adds a renderer to the draw stack.
func (g *Game) AddRenderer(renderer ...engine.Renderer) {
	g.renderers = append(g.renderers, renderer...)
}

// Layout converts window size to screen size, using the Game's LayoutFunc.
//
// Some constraints are applied to work around ebiten quirks/bugs.
func (g *Game) Layout(ow, oh int) (int, int) {
	/*
		// no resize happened
		if ow == g.ow && oh == g.oh {
			return ow, oh
		}

		g.ow, g.oh = ow, oh
	*/

	w, h := g.layoutFunc(ow, oh, g.vw, g.vh)

	if w < 1 {
		w = 1
	}

	if h < 1 {
		h = 1
	}

	// to avoid a panic in ebiten related to opengl
	const openglMax = 16384

	if w > openglMax {
		w = openglMax
	}

	if h > openglMax {
		h = openglMax
	}

	g.w, g.h = w, h

	return w, h
}

// Update runs the tick functions.
func (g *Game) Update() error {
	g.tickFunc()

	for _, renderer := range g.renderers {
		renderer.Tick()
	}

	return nil
}

// Draw runs the draw functions.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, renderer := range g.renderers {
		log.Println("set viewport", g.w, g.h)
		renderer.SetViewport(g.w, g.h)

		switch r := renderer.(type) {
		case *Renderer:
			r.draw(screen)
		case *IsoRenderer:
			r.draw(screen)
		}
	}
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
