package engine

import "image"

type Context struct {
	Renderer
	Collider

	entities map[string][]Entity
}

func NewContext(renderer Renderer, collider Collider) *Context {
	return &Context{
		Renderer: renderer,
		Collider: collider,
		entities: make(map[string][]Entity),
	}
}

func (c *Context) AddEntity(entities ...Entity) {
	for _, e := range entities {
		if c.Collider != nil {
			e.SetCollider(c.Collider)
		}
		c.AddImage(e.Images()...)

		c.entities[e.Class()] = append(
			c.entities[e.Class()],
			e,
		)
	}
}

func (c *Context) Tick() {
	for class, entities := range c.entities {
		var i int
		for _, e := range entities {
			if e.IsDisposed() {
				continue
			}

			// TODO make this good
			// spacial partitioning, etc
			if hitbox, ok := e.(Hitbox); ok {
				for _, targetClass := range hitbox.HitClasses() {
					for _, targetEntity := range c.entities[targetClass] {
						targetHitbox, ok := targetEntity.(Hitbox)
						if !ok {
							continue
						}

						sourcePos := hitbox.Position()
						sourceBounds := hitbox.HitBounds()
						targetPos := targetHitbox.Position()
						targetBounds := targetHitbox.HitBounds()

						if image.Rect(
							int(sourcePos.X)+sourceBounds.Min.X,
							int(sourcePos.Y)+sourceBounds.Min.Y,
							int(sourcePos.X)+sourceBounds.Max.X,
							int(sourcePos.Y)+sourceBounds.Max.Y,
						).Overlaps(
							image.Rect(
								int(targetPos.X)+targetBounds.Min.X,
								int(targetPos.Y)+targetBounds.Min.Y,
								int(targetPos.X)+targetBounds.Max.X,
								int(targetPos.Y)+targetBounds.Max.Y,
							),
						) {
							hitbox.Hit(targetEntity)
						}
					}
				}
			}

			e.Tick()
			c.entities[class][i] = e
			i++
		}

		for j := i; j < len(c.entities[class]); j++ {
			c.entities[class][j] = nil
		}
		c.entities[class] = c.entities[class][:i]
	}
}
