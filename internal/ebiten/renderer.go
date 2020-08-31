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

// RemoveImage removes images from the draw stack.
func (r *Renderer) RemoveImage(images ...engine.Image) {
	for _, remove := range images {
		for i, img := range r.imgs {
			if img == remove {
				r.imgs[i] = r.imgs[len(r.imgs)-1]
				r.imgs[len(r.imgs)-1] = nil
				r.imgs = r.imgs[:len(r.imgs)-1]
				break
			}
		}
	}
}

// draw renders all images in the draw stack.
func (r *Renderer) draw(screen *ebiten.Image) {
	for _, img := range r.imgs {
		eimg := img.(*Image)
		op := new(ebiten.DrawImageOptions)

		op.GeoM.Translate(eimg.tx, eimg.ty)
		op.GeoM.Scale(eimg.sx, eimg.sy)
		op.GeoM.Rotate(eimg.d)

		screen.DrawImage(eimg.img, op)
	}
}
