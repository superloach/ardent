package engine

import "image"

type Component interface {
	NewImageFromPath(string) (Image, error)
	NewImageFromImage(image.Image) Image

	NewRenderer() Renderer
}
