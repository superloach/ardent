package engine

type Context struct {
	Renderer
	Collider

	entities []Entity
}

func (c *Context) AddEntity(entities ...Entity) {
	for _, e := range entities {
		if c.Collider != nil {
			e.SetCollider(c.Collider)
		}
		c.Renderer.AddImage(e.Images()...)
	}

	c.entities = append(c.entities, entities...)
}

func (c *Context) Tick() {
	for _, e := range c.entities {
		// TODO check if disposed
		e.Tick()
	}
}
