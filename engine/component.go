package engine

import "image"

type Component interface {
	NewAssetFromPath(string) (Asset, error)

	NewImageFromPath(string) (Image, error)
	NewImageFromAssetPath(string) (Image, error)
	NewImageFromImage(image.Image) Image

	NewAtlasFromAssetPath(string) (Atlas, error)

	NewRenderer() Renderer
}
