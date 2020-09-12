package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type Asset struct {
	img   *Image
	atlas *Atlas
}

func (a *Asset) ToImage() engine.Image {
	return a.img
}

func (a *Asset) ToAtlas() engine.Atlas {
	return a.atlas
}

func (a *Asset) UnmarshalBinary(data []byte) error {
	ca := common.NewAsset()
	if err := ca.UnmarshalBinary(data); err != nil {
		return err
	}

	switch ca.Type {
	case common.AssetTypeImage:
		img, _ := ebiten.NewImageFromImage(ca.Img, ebiten.FilterDefault)
		a.img = &Image{
			img: img,
			sx:  1,
			sy:  1,
		}
	case common.AssetTypeAtlas:
		img, _ := ebiten.NewImageFromImage(ca.Img, ebiten.FilterDefault)
		a.atlas = &Atlas{
			img:     img,
			regions: ca.AtlasMap,
			cache:   make(map[string]engine.Image),
		}
	case common.AssetTypeAnimation:
	case common.AssetTypeSound:
	default:
		panic("Invalid asset type")
	}

	return nil
}
