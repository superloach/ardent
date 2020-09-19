package headless

import "github.com/split-cube-studios/ardent/engine"

type IsoRenderer struct{}

func (i IsoRenderer) SetTilemap(tilemap engine.Tilemap) {}

func (i IsoRenderer) SetCamera(camera engine.Camera) {}

func (i IsoRenderer) AddImage(image ...engine.Image) {}
