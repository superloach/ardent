package ebiten

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
)

type component struct{}

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

func (c component) NewImageFromImage(img image.Image) engine.Image {
	eimg, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return &Image{img: eimg, op: new(ebiten.DrawImageOptions)}
}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}
