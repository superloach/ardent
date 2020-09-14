package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

type Asset struct{}

func (a Asset) ToImage() engine.Image {
	return new(Image)
}

func (a Asset) ToAtlas() engine.Atlas {
	return new(Atlas)
}

func (a Asset) ToAnimation() engine.Animation {
	return new(Animation)
}
