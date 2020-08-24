package headless

import (
	"image"

	"github.com/split-cube-studios/ardent/engine"
)

type component struct{}

func (c component) NewImageFromPath(path string) (engine.Image, error) {
	return new(Image), nil
}

func (c component) NewImageFromImage(img image.Image) engine.Image {
	return new(Image)
}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}
