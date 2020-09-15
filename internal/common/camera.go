package common

type Camera struct {
	x, y float64
}

func (c *Camera) LookAt(x, y float64) {
	c.x, c.y = x, y
}

func (c *Camera) Position() (float64, float64) {
	return c.x, c.y
}
