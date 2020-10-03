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

func (t *Tilemap) isoToIndex(x, y float64) (int, int) {
	/*
		i := y*4/float64(t.Width) + 1
		j := (i*float64(t.Width) + 2*x) / float64(2*t.Width)

		//	return int(math.Round(i-j)) + 1, int(math.Round(j))
		return int(math.Round(j)), int(math.Round(i-j)) + 1
	*/

	ix := int(math.Ceil((x/float64(t.Width/2) + y/float64(t.Width/4)) / 2))
	iy := int(math.Ceil((y/float64(t.Width/4) - x/float64(t.Width/2)) / 2))

	return ix + 1, iy + 1
}

func (t *Tilemap) indexToIso(i, j int) (float64, float64) {
	/*
		i += j

		y := i * t.Width / 4
		x := ((j - i/2) * t.Width) - (t.Width * (i % 2) / 2)

		return float64(x), float64(y)
	*/

	x := (i - j) * (t.Width / 2)
	y := (i + j) * (t.Width / 4)

	return float64(x), float64(y)
}

func (t *Tilemap) getTileValue(x, y, z int) int {
	if z >= len(t.Data) || x >= len(t.Data[0]) || y >= len(t.Data[0][0]) ||
		z < 0 || x < 0 || y < 0 {
		return 0
	}
	return t.Data[z][y][x]
}
