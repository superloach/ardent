package engine

type Asset interface {
	ToImage() (Image, error)
	ToAtlas() (Atlas, error)
}
