package engine

// A Renderer is a basic context for drawing images.
type Renderer interface {
	// AddImage adds one or more images
	// to the renderer's draw stack.
	// Images are drawn in the order they are added.
	AddImage(...Image)

	// Remove image removes one or more images
	// from the renderer's draw stack.
	RemoveImage(...Image)
}
