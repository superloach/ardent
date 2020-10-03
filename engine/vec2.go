package engine

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) AngleTo(v2 Vec2) float64 {
	return math.Atan2(v2.Y-v.Y, v2.X-v.X)
}
