package common

import "github.com/split-cube-studios/ardent/engine"

type Tilemap struct {
	Width  int
	Data   [2][][]int
	Mapper map[int]engine.Image
}
