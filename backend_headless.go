// +build headless

package ardent

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/headless"
)

func newGame(
	title string,
	w, h int,
	flags byte,
	tickFunc func(),
	layout engine.LayoutHandler,
) engine.Game {
	return headless.NewGame(tickFunc)
}
