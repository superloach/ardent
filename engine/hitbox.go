package engine

import "image"

type Hitbox interface {
	Position() Vec2
	HitBounds() image.Rectangle // TODO upgrade to float based approach?
	HitClasses() []string
	Hit(Entity)
}
