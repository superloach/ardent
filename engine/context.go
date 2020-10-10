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
	var i int
	for _, e := range c.entities {
		if e.IsDisposed() {
			continue
		}

		e.Tick()
		c.entities[i] = e
		i++
	}

	for j := i; j < len(c.entities); j++ {
		c.entities[j] = nil
	}
	c.entities = c.entities[:i]
}
