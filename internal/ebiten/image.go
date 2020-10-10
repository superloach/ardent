package ebiten

import "github.com/hajimehoshi/ebiten"

// Image is an ebiten implementation of engine.Image
type Image struct {
	img *ebiten.Image

	tx, ty float64
	sx, sy float64
	d      float64

	disposed bool
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
func (i *Image) Size() (int, int) {
	return i.img.Size()
}

// Dispose marks the image to be disposed.
func (i *Image) Dispose() {
	i.disposed = true
}

// IsDisposed indicates if the image has been dispoed.
func (i *Image) IsDisposed() bool {
	return i.disposed
}
