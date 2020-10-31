package ebiten

import (
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type IsoRenderer struct {
	Renderer

	tilemap *common.Tilemap
}

type isoRendererImage struct {
	img        *Image
	isTile     bool
	tileHeight int
}

func (r *IsoRenderer) SetTilemap(tilemap engine.Tilemap) {
	r.tilemap = tilemap.(*common.Tilemap)
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
				x -= float64(tw / 2)
				y -= float64(tw)

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
						img:        img.(*Image).img,
						tx:         x,
						ty:         y,
						sx:         1,
						sy:         1,
						r:          1,
						g:          1,
						b:          1,
						alpha:      1,
						renderable: true,
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

	for _, img := range r.imgs {
		if !img.IsRenderable() {
			continue
		}

		var tmpImage *isoRendererImage

		var x, y float64

		switch img.(type) {
		case *Image:
			i := img.(*Image)
			w, h := i.Size()
			tmpImage = &isoRendererImage{
				img: &Image{
					img:               i.img,
					tx:                i.tx - i.originX*float64(w),
					ty:                i.ty - i.originY*float64(h),
					ox:                i.ox,
					oy:                i.oy,
					sx:                i.sx,
					sy:                i.sy,
					originX:           i.originX,
					originY:           i.originY,
					d:                 i.d,
					z:                 i.z,
					r:                 i.r,
					g:                 i.g,
					b:                 i.b,
					alpha:             i.alpha,
					renderable:        i.renderable,
					roundTranslations: i.roundTranslations,
				},
			}

			x, y = i.tx, i.ty

		case *Animation:
			a := img.(*Animation)
			w, h := a.Size()
			tmpImage = &isoRendererImage{
				img: &Image{
					img:               a.getFrame(),
					tx:                a.tx - a.originX*float64(w),
					ty:                a.ty - a.originY*float64(h),
					ox:                a.ox,
					oy:                a.oy,
					sx:                a.sx,
					sy:                a.sy,
					originX:           a.originX,
					originY:           a.originY,
					d:                 a.d,
					z:                 a.z,
					r:                 a.r,
					g:                 a.g,
					b:                 a.b,
					alpha:             a.alpha,
					renderable:        a.renderable,
					roundTranslations: a.roundTranslations,
				},
			}

			x, y = a.tx, a.ty

		default:
			panic("Invalid image type")
		}

		if x-cx < -800 || x-cx > float64(r.w+800) {
			continue
		}

		if y-cy < -800 || y-cy > float64(r.h+800) {
			continue
		}

		layers[len(layers)-1] = append(layers[len(layers)-1], tmpImage)
	}

	for _, layer := range layers {
		sort.SliceStable(layer, func(i, j int) bool {
			var ty1, ty2 float64

			img := layer[i].img
			_, h := img.Size()

			if layer[i].isTile {
				ty1 = img.ty + float64(h-layer[i].tileHeight/8)
			} else {
				ty1 = img.ty + float64(h)
			}

			img = layer[j].img
			_, h = img.Size()

			if layer[j].isTile {
				ty2 = img.ty + float64(h-layer[j].tileHeight/8)
			} else {
				ty2 = img.ty + float64(h)
			}

			return ty1 < ty2
		})

		sort.SliceStable(layer, func(i, j int) bool {
			return layer[i].img.z < layer[j].img.z
		})

		for _, isoImage := range layer {
			img := isoImage.img

			op := new(ebiten.DrawImageOptions)
			w, h := img.Size()

			op.GeoM.Scale(img.sx, img.sy)
			op.GeoM.Translate(
				-img.originX*float64(w),
				-img.originY*float64(h),
			)
			op.GeoM.Rotate(img.d)
			op.GeoM.Translate(
				img.originX*float64(w),
				img.originY*float64(h),
			)

			x, y := img.tx+img.ox, img.ty+img.oy
			if img.roundTranslations {
				x, y = math.Round(x), math.Round(y)
			}

			op.GeoM.Translate(
				x-cx,
				y-cy,
			)

			op.ColorM.Scale(img.r, img.g, img.b, img.alpha)

			screen.DrawImage(img.img, op)
		}
	}
}
