package common

import (
	"math"

	"github.com/split-cube-studios/ardent/engine"
)

type Tilemap struct {
	Width  int
	Data   [2][][]int
	Mapper map[int]engine.Image
}

func (t *Tilemap) IsoToIndex(x, y float64) (int, int) {
	ix := int(math.Ceil((x/float64(t.Width/2) + y/float64(t.Width/4)) / 2))
	iy := int(math.Ceil((y/float64(t.Width/4) - x/float64(t.Width/2)) / 2))

	return ix + 1, iy + 1
}

func (t *Tilemap) IndexToIso(i, j int) (float64, float64) {
	x := (i - j) * (t.Width / 2)
	y := (i + j) * (t.Width / 4)

	return float64(x), float64(y)
}

func (t *Tilemap) GetTileValue(x, y, z int) int {
	if z >= len(t.Data) || x >= len(t.Data[0]) || y >= len(t.Data[0][0]) ||
		z < 0 || x < 0 || y < 0 {
		return 0
	}
	return t.Data[z][y][x]
}
