package common

import (
	"github.com/split-cube-studios/ardent/engine"
)

type Collider struct {
	m *Tilemap
}

func (c *Collider) SetTilemap(m engine.Tilemap) {
	c.m = m.(*Tilemap)
}

func (c *Collider) Resolve(x, y float64) (float64, float64) {
	if c.m == nil {
		return x, y
	}

	ix, iy := c.m.isoToIndex(x, y)
	if c.m.getTileValue(ix, iy, 1) == 0 {
		return x, y
	}

	tileX, tileY := c.m.indexToIso(ix, iy)
	centerX, centerY := tileX+float64(c.m.Width/2), tileY+float64(c.m.Width/4)

	a1 := centerY - y
	b1 := x - centerX
	c1 := a1*x + b1*y

	ax, ay := tileX, centerY // left corner
	bx, by := centerX, tileY // top corner

	// right corner
	if x > centerX {
		ax += float64(c.m.Width)
	}

	// bottom corner
	if y > centerY {
		by += float64(c.m.Width / 2)
	}

	a2 := by - ay
	b2 := ax - bx
	c2 := a2*ax + b2*ay

	d := a1*b2 - a2*b1
	if d == 0 {
		return x, y
	}

	return (b2*c1 - b1*c2) / d, (a1*c2 - a2*c1) / d
}
