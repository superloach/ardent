package ebiten

import "github.com/split-cube-studios/ardent/engine"

type Atlas struct {
	img *Image
}

func (a *Atlas) GetImage(k string) engine.Image {
	return nil
}
