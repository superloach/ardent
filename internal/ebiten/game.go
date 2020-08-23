package ebiten

import "github.com/hajimehoshi/ebiten"

// Game is an ebiten implementation
// of engine.Game.
type Game struct {
	tickFunc   func()
	drawFunc   func()
	layoutFunc func(int, int) (int, int)
}

// NewGame returns an instantiated game.
func NewGame(
	tickFunc func(),
	drawFunc func(),
	layoutFunc func(int, int) (int, int),
) *Game {
	return &Game{
		tickFunc:   tickFunc,
		drawFunc:   drawFunc,
		layoutFunc: layoutFunc,
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

// Layout is called when the window resizes.
func (g Game) Layout(ow, oh int) (int, int) {
	return g.layoutFunc(ow, oh)
}

// Update runs the tick and draw functions.
func (g *Game) Update(screen *ebiten.Image) error {
	g.tickFunc()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.drawFunc()

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
