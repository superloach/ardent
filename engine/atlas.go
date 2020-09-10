package engine

type Atlas interface {
	GetImage(string) engine.Image
}
