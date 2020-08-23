package ardent

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/ebiten"
)

// Backend flag type
type Backend byte

// Backend options
const (
	EBITEN Backend = 1 << iota
)

// NewGame creates a new game instance for a given backend.
func NewGame(
	backend Backend,
	tickFunc func(),
	drawFunc func(),
	layoutFunc func(int, int) (int, int),
) engine.Game {
	switch backend {
	case EBITEN:
		return ebiten.NewGame(
			tickFunc,
			drawFunc,
			layoutFunc,
		)
	default:
		panic("Invalid backend")
	}
}
