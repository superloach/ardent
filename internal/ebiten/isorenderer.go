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

func (i *IsoRenderer) SetTilemap(tilemap engine.Tilemap) {
	i.tilemap = tilemap.(*common.Tilemap)
}

func (i *IsoRenderer) SetCamera(camera engine.Camera) {
	i.camera = camera
}

func (i *IsoRenderer) AddImage(images ...engine.Image) {
	i.images = append(i.images, images...)
}

func (i *IsoRenderer) tick() {
	for _, img := range i.images {
		anim, ok := img.(*Animation)
		if ok {
			anim.tick()
		}
	}
}

func (i *IsoRenderer) tilemapToIsoLayers() [][]*isoRendererImage {
	if i.tilemap == nil {
		return make([][]*isoRendererImage, 1)
	}

	tw := i.tilemap.Width
	data := i.tilemap.Data
	mapper := i.tilemap.Mapper

	w, h := len(data[0][0]), len(data[0])

	layers := make([][]*isoRendererImage, len(data))

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

				layers[i] = append(layers[i], &isoRendererImage{
					img: &Image{
						img: img.(*Image).img,
						tx:  float64(x),
						ty:  float64(y),
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

func (i *IsoRenderer) draw(screen *ebiten.Image) {
	var cx, cy float64
	if i.camera != nil {
		cx, cy = i.camera.Position()
		cx, cy = cx-float64(i.w/2), cy-float64(i.h/2)
	}

	layers := i.tilemapToIsoLayers()

	for _, img := range i.images {
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

func (i *IsoRenderer) setViewport(w, h int) {
	i.w, i.h = w, h
}
