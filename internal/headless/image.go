package headless

// Image is a headless implementation of engine.Image
type Image struct{}

// Translate sets the image translation.
func (i Image) Translate(x float64, y float64) {
	// NOOP
}

// Scale sets the image scale.
func (i Image) Scale(x float64, y float64) {
	// NOOP
}

// Rotate sets the image rotation.
func (i Image) Rotate(d float64) {
	// NOOP
}

// Size returns the image size.
func (i Image) Size() (int, int) {
	return 0, 0
}

func (i Image) Dispose() {}

func (i Image) IsDisposed() bool {
	return false
}
