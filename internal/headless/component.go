package headless

import (
	"image"

	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
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

func (c component) NewAnimationFromAssetPath(path string) (engine.Animation, error) {
	return new(Animation), nil
}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}

func (c component) NewIsoRenderer() engine.IsoRenderer {
	return new(IsoRenderer)
}

func (c component) NewTilemap(width int, data [2][][]int, mapper map[int]engine.Image) engine.Tilemap {
	return new(common.Tilemap)
}

func (c component) NewCamera() engine.Camera {
	return new(common.Camera)
}
