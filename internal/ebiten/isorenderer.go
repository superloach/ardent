package ebiten

import (
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type IsoRenderer struct {
	tilemap *common.Tilemap
	camera  engine.Camera
	images  []engine.Image

	w, h int
}

type isoRendererImage struct {
	img        *Image
	isTile     bool
	tileHeight int
}

func (r *IsoRenderer) SetTilemap(tilemap engine.Tilemap) {
	r.tilemap = tilemap.(*common.Tilemap)
}

func (r *IsoRenderer) SetCamera(camera engine.Camera) {
	r.camera = camera
}

func (r *IsoRenderer) AddImage(images ...engine.Image) {
	r.images = append(r.images, images...)
}

func (r *IsoRenderer) Tick() {
	for _, img := range r.images {
		anim, ok := img.(*Animation)
		if ok {
			anim.tick()
		}
	}
}

func (r *IsoRenderer) tilemapToIsoLayers(cx, cy float64) [][]*isoRendererImage {
	if r.tilemap == nil {
		return make([][]*isoRendererImage, 1)
	}

	tw := r.tilemap.Width
	data := r.tilemap.Data
	mapper := r.tilemap.Mapper

	layers := make([][]*isoRendererImage, len(data))

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			for k := 0; k < len(data[i][j]); k++ {
				x, y := r.tilemap.IndexToIso(j, k)

				if x-cx < -800 || x-cx > float64(r.w+800) {
					continue
				}

				if y-cy < -800 || y-cy > float64(r.h+800) {
					continue
				}

				img := mapper[data[i][k][j]]
				if img == nil {
					continue
				}

				if i != 0 {
					_, ih := img.Size()
					y -= float64(ih - tw/4)
				}

				layers[i] = append(layers[i], &isoRendererImage{
					img: &Image{
						img: img.(*Image).img,
						tx:  x,
						ty:  y,
						sx:  1,
						sy:  1,
					},
					isTile:     true,
					tileHeight: tw / 2,
				})
			}
		}
	}

	return layers
}

func (r *IsoRenderer) draw(screen *ebiten.Image) {
	var cx, cy float64
	if r.camera != nil {
		cx, cy = r.camera.Position()
		cx, cy = cx-float64(r.w/2), cy-float64(r.h/2)
	}

	layers := r.tilemapToIsoLayers(cx, cy)

	for _, img := range r.images {
		var tmpImage *isoRendererImage

		switch img.(type) {
		case *Image:
			tmpImage = &isoRendererImage{
				img: img.(*Image),
			}

		case *Animation:
			a := img.(*Animation)
			tmpImage = &isoRendererImage{
				img: &Image{
					img: a.getFrame(),
					tx:  a.tx,
					ty:  a.ty,
					sx:  a.sx,
					sy:  a.sy,
					d:   a.d,
				},
			}

		default:
			panic("Invalid image type")
		}

		layers[len(layers)-1] = append(layers[len(layers)-1], tmpImage)
	}

	for _, layer := range layers {
		sort.Slice(layer, func(i, j int) bool {
			var ty1, ty2 float64

			img := layer[i].img
			_, h := img.Size()

			if layer[i].isTile {
				ty1 = img.ty + float64(h-layer[i].tileHeight/2)
			} else {
				ty1 = img.ty + float64(h/2)
			}

			img = layer[j].img
			_, h = img.Size()

			if layer[j].isTile {
				ty2 = img.ty + float64(h-layer[j].tileHeight/2)
			} else {
				ty2 = img.ty + float64(h/2)
			}

			return ty1 < ty2
		})

		for _, isoImage := range layer {
			img := isoImage.img

			op := new(ebiten.DrawImageOptions)
			op.GeoM.Translate(img.tx-cx, img.ty-cy)
			op.GeoM.Scale(img.sx, img.sy)
			op.GeoM.Rotate(img.d)

			screen.DrawImage(img.img, op)
		}
	}
}

func (r *IsoRenderer) setViewport(w, h int) {
	r.w, r.h = w, h
}
