package engine

type Collider interface {
	SetTilemap(Tilemap)
	Resolve(float64, float64) (float64, float64)
}
