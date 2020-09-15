package ebiten

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type component struct{}

func (c component) NewAssetFromPath(path string) (engine.Asset, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open asset path: %w", err)
	}
	defer f.Close()

	a := new(Asset)
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode asset: %w", err)
	}

	return a, a.UnmarshalBinary(d)
}

func (c component) NewImageFromPath(path string) (engine.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open image path: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode image: %w", err)
	}

	return c.NewImageFromImage(img), nil
}

func (c component) NewImageFromAssetPath(path string) (engine.Image, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	return a.ToImage(), nil
}

func (c component) NewImageFromImage(img image.Image) engine.Image {
	eimg, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return &Image{
		img: eimg,
		sx:  1,
		sy:  1,
	}
}

func (c component) NewAtlasFromAssetPath(path string) (engine.Atlas, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	return a.ToAtlas(), nil
}

func (c component) NewAnimationFromAssetPath(path string) (engine.Animation, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	return a.ToAnimation(), nil
}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}

func (c component) NewIsoRenderer() engine.IsoRenderer {
	return new(IsoRenderer)
}

func (c component) NewTilemap(width int, data [2][][]int, mapper map[int]engine.Image) engine.Tilemap {
	return &common.Tilemap{
		Width:  width,
		Data:   data,
		Mapper: mapper,
	}
}

func (c component) NewCamera() engine.Camera {
	return new(common.Camera)
}
