package ardent

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/ebiten"
	"github.com/split-cube-studios/ardent/internal/headless"
)

// Backend flag type
type Backend byte

// Backend options
const (
	HEADLESS Backend = 1 << iota
	EBITEN
)

// NewGame creates a new game instance for a given backend.
func NewGame(
	title string,
	w, h int,
	flags byte,
	backend Backend,
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) engine.Game {
	switch backend {
	case HEADLESS:
		return headless.NewGame(
			tickFunc,
			nil,
		)
	case EBITEN:
		return ebiten.NewGame(
			title,
			w,
			h,
			flags,
			tickFunc,
			layoutFunc,
		)
	default:
		panic("Invalid backend")
	}
}
