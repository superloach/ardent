package engine

type Component interface {
	NewRenderer() Renderer
}
