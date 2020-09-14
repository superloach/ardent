package engine

type Asset interface {
	ToImage() Image
	ToAtlas() Atlas
	ToAnimation() Animation
}
