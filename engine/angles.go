package engine

import "math"

// Cardinal directions
const (
	N = 1 << iota
	E
	S
	W

	NE = N | E
	NW = N | W
	SE = S | E
	SW = S | W
)

const diag = 3.435 * math.Pi / 180

// CardinalToAngle is a map of cardinal directions
// to angles in dimetric space.
var CardinalToAngle = map[byte]float64{
	N:  3 * math.Pi / 2,            // 90
	E:  0,                          // 0
	S:  math.Pi / 2,                // 270
	W:  math.Pi,                    // 180
	SE: math.Pi/6 - diag,           // 26.565
	SW: 5*math.Pi/6 + diag,         // 153.435
	NE: -(math.Pi / 6) + diag,      // 333.435
	NW: math.Pi + math.Pi/6 - diag, // 206.565,
}

func AngleToCardinal(angle float64) byte {
	var cardinal byte
	closest := math.MaxFloat64

	for k, v := range CardinalToAngle {
		diff := math.Abs(angle - v)
		if diff < math.Abs(closest) {
			cardinal = k
			closest = diff
		}
	}

	return cardinal
}
