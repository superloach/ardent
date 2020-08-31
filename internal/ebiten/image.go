package ebiten

import "github.com/hajimehoshi/ebiten"

// Image is an ebiten implementation of engine.Image
type Image struct {
	img *ebiten.Image

	tx, ty float64
	sx, sy float64
	d      float64
}

// Translate sets the image translation.
func (i *Image) Translate(x, y float64) {
	i.tx, i.ty = x, y
}

// Scale sets the image scale.
func (i *Image) Scale(x, y float64) {
	i.sx, i.sy = x, y
}

// Rotate sets the image rotation.
func (i *Image) Rotate(d float64) {
	i.d = d
}

// Size returns the image size.
func (i Image) Size() (int, int) {
	return i.img.Size()
}
