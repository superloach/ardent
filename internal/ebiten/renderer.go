package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
)

// Renderer is a simple ebiten renderer.
type Renderer struct {
	imgs []engine.Image
}

// AddImage adds images to the draw stack.
func (r *Renderer) AddImage(images ...engine.Image) {
	r.imgs = append(r.imgs, images...)
}

// draw renders all images in the draw stack.
func (r *Renderer) draw(screen *ebiten.Image) {
	for _, img := range r.imgs {
		screen.DrawImage(img.(*Image).img, img.(*Image).op)
	}

	r.purgeBuffer(false)
}

// purgeBuffer clears the image buffer.
// A bool can be passed to optionally reallocate
// the buffer, or to simply reslice it.
func (r *Renderer) purgeBuffer(realloc bool) {
	if realloc {
		r.imgs = make([]engine.Image, 0)
		return
	}
	r.imgs = r.imgs[:0]
}
