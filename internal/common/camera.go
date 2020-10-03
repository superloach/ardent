package common

import "github.com/split-cube-studios/ardent/engine"

type Camera struct {
	engine.Vec2
}

func (c *Camera) LookAt(x, y, t float64) {
	c.Vec2 = c.Vec2.Lerp(engine.Vec2{x, y}, t)
}

func (c *Camera) Position() (float64, float64) {
	return c.X, c.Y
}
