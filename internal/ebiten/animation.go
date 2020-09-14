package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/internal/common"
)

type Animation struct {
	img   *ebiten.Image
	state string

	w, h                     uint16
	fpsCounter, frameCounter uint16

	anims map[string]common.Animation
	cache map[uint16]*ebiten.Image

	tx, ty float64
	sx, sy float64
	d      float64
}

func (a *Animation) SetState(state string) {
	if a.state == state {
		return
	}

	a.state = state
	a.fpsCounter, a.frameCounter = 0, 0
}

func (a *Animation) tick() {
	fps := a.anims[a.state].Fps
	if a.fpsCounter == 60/fps {
		a.frameCounter++
		a.fpsCounter = 0
		return
	}

	a.fpsCounter++
}

func (a *Animation) Translate(x, y float64) {
	a.tx, a.ty = x, y
}

func (a *Animation) Scale(x, y float64) {
	a.sx, a.sy = x, y
}

func (a *Animation) Rotate(d float64) {
	a.d = d
}

func (a *Animation) Size() (int, int) {
	return int(a.w), int(a.h)
}

func (a *Animation) getFrame() *ebiten.Image {
	anim, ok := a.anims[a.state]
	if !ok {
		return nil
	}

	frameKey := (a.frameCounter % (anim.End - anim.Start)) + anim.Start
	frame, ok := a.cache[frameKey]
	if ok {
		return frame
	}

	w, _ := a.img.Size()
	xtiles := uint16(w) / a.w
	x := (frameKey % xtiles) * a.w
	y := (frameKey / xtiles) * a.h

	img := a.img.SubImage(
		image.Rect(
			int(x),
			int(y),
			int(x+a.w),
			int(y+a.h),
		),
	).(*ebiten.Image)
	a.cache[frameKey] = img

	return img
}
