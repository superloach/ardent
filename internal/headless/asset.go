package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

type Asset struct{}

func (a Asset) ToImage() (engine.Image, error) {
	return new(Image), nil
}

func (a Asset) ToAtlas() (engine.Atlas, error) {
	return new(Atlas), nil
}
