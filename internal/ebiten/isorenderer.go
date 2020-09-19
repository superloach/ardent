package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type IsoRenderer struct {
	tilemap *common.Tilemap
	camera  engine.Camera

	w, h int
}

func (i *IsoRenderer) SetTilemap(tilemap engine.Tilemap) {
	i.tilemap = tilemap.(*common.Tilemap)
}

func (i *IsoRenderer) SetCamera(camera engine.Camera) {
	i.camera = camera
}

func (i *IsoRenderer) draw(screen *ebiten.Image) {
	if i.tilemap == nil {
		return
	}

	tw := i.tilemap.Width
	data := i.tilemap.Data
	mapper := i.tilemap.Mapper

	w, h := len(data[0][0]), len(data[0])

	var cx, cy float64

	if i.camera != nil {
		cx, cy = i.camera.Position()
		cx, cy = cx-float64(i.w/2), cy-float64(i.h/2)
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < w*h; j++ {
			y := j * tw / 4

			for k := 0; k <= j; k++ {
				if !(j-k < w && k < h) {
					continue
				}

				x := ((k - j/2) * tw) - (tw * (j % 2) / 2)

				img := mapper[data[i][j-k][k]]
				if img == nil {
					continue
				}

				y := y // shadow var for modification
				if i != 0 {
					_, ih := img.Size()
					y -= ih - tw/4
				}

				op := new(ebiten.DrawImageOptions)
				op.GeoM.Translate(float64(x)-cx, float64(y)-cy)

				screen.DrawImage(img.(*Image).img, op)
			}
		}
	}
}

func (i *IsoRenderer) setViewport(w, h int) {
	i.w, i.h = w, h
}
