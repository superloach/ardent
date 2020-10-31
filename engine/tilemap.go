package engine

type Tilemap interface{}

type TileOverlapEvent func(bool, Image, interface{}) interface{}
