package render

import "github.com/split-cube-studios/ardent/engine"

type BasicRenderer struct {
	imgs []engine.Image
}

func (r *BasicRenderer) AddImage(images ...engine.Image) {
	r.imgs = append(r.imgs, images)
}

func (r *BasicRenderer) Draw() {

	purgeBuffer(false)
}

// purgeBuffer clears the image buffer.
// A bool can be passed to optionally reallocate
// the buffer, or to simply reslice it.
func (r *BasicRenderer) purgeBuffer(realloc bool) {
	if realloc {
		r.imgs = make([]engine.Image, 0)
		return
	}
	r.imgs = r.imgs[:0]
}
