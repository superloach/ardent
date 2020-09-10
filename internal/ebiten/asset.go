package ebiten

import (
	"github.com/split-cube-studios/ardent/engine"
)

type Asset struct{}

func (a *Asset) ToImage() (engine.Image, error) {
	return nil, nil
}

func (a *Asset) ToAtlas() (engine.Atlas, error) {
	return nil, nil
}
