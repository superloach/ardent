package headless

import (
	"image"

	"github.com/split-cube-studios/ardent/engine"
)

type component struct{}

func (c component) NewAssetFromPath(path string) (engine.Asset, error) {
	return new(Asset), nil
}

func (c component) NewImageFromPath(path string) (engine.Image, error) {
	return new(Image), nil
}

func (c component) NewImageFromAssetPath(path string) (engine.Image, error) {
	return new(Image), nil
}

func (c component) NewImageFromImage(img image.Image) engine.Image {
	return new(Image)
}

func (c component) NewAtlasFromAssetPath(path string) (engine.Atlas, error) {
	return new(Atlas), nil
}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}
