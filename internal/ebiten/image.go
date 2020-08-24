package ebiten

import "github.com/hajimehoshi/ebiten"

// Image is an ebiten implementation of engine.Image
type Image struct {
	img *ebiten.Image
	op  *ebiten.DrawImageOptions
}

// Translate sets the image translation.
func (i *Image) Translate(x float64, y float64) {
	i.op.GeoM.Translate(x, y)
}

// Scale sets the image scale.
func (i *Image) Scale(x float64, y float64) {
	i.op.GeoM.Scale(x, y)
}

// Rotate sets the image rotation.
func (i *Image) Rotate(d float64) {
	i.op.GeoM.Rotate(d)
}

// Size returns the image size.
func (i Image) Size() (int, int) {
	return i.img.Size()
}
