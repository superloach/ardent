package engine

type Collider interface {
	SetTilemap(Tilemap)
	Resolve(Vec2, Vec2) Vec2
}
