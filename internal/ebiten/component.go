package ebiten

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/ebiten/render"
)

type component struct{}

func (c component) NewRenderer() engine.Renderer {
	return new(render.Renderer)
}
