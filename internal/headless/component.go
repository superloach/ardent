package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

type component struct{}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}
