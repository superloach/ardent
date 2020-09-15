package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
)

// Renderer is a simple ebiten renderer.
type Renderer struct {
	imgs   []engine.Image
	camera engine.Camera
}

// AddImage adds images to the draw stack.
func (r *Renderer) AddImage(images ...engine.Image) {
	r.imgs = append(r.imgs, images...)
}

func (r *Renderer) SetCamera(camera engine.Camera) {
	r.camera = camera
}

func (r *Renderer) tick() {
	for _, img := range r.imgs {
		anim, ok := img.(*Animation)
		if ok {
			anim.tick()
		}
	}
}

// draw renders all images in the draw stack.
func (r *Renderer) draw(screen *ebiten.Image) {
	var (
		eimg   *ebiten.Image
		tx, ty float64
		sx, sy float64
		d      float64
		cx, cy float64
	)

	if r.camera != nil {
		cx, cy = r.camera.Position()
	}

	for _, img := range r.imgs {
		switch img.(type) {
		case *Image:
			i := img.(*Image)
			eimg = i.img
			tx, ty = i.tx, i.ty
			sx, sy = i.sx, i.sy
			d = i.d

		case *Animation:
			a := img.(*Animation)
			eimg = a.getFrame()
			tx, ty = a.tx, a.ty
			sx, sy = a.sx, a.sy
			d = a.d

		default:
			panic("Invalid image type")
		}

		op := new(ebiten.DrawImageOptions)

		op.GeoM.Translate(tx-cx, ty-cy)
		op.GeoM.Scale(sx, sy)
		op.GeoM.Rotate(d)

		screen.DrawImage(eimg, op)
	}
}
