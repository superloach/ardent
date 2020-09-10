package headless

import "github.com/split-cube-studios/ardent/engine"

type Atlas struct{}

func (a Atlas) GetImage(k string) engine.Image {
	return new(Image)
}
