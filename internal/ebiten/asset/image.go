package asset

import "github.com/hajimehoshi/ebiten"

// Image is an ebiten implementation of engine.Image
type Image struct {
	Img *ebiten.Image
	Op  *ebiten.DrawImageOptions
}

// Translate sets the image translation.
func (i *Image) Translate(x float64, y float64) {
	i.Op.GeoM.Translate(x, y)
}

// Scale sets the image scale.
func (i *Image) Scale(x float64, y float64) {
	i.Op.GeoM.Scale(x, y)
}

// Rotate sets the image rotation.
func (i *Image) Rotate(d float64) {
	i.Op.GeoM.Rotate(d)
}

// Size returns the image size.
func (i Image) Size() (int, int) {
	return i.Img.Size()
}
