package engine

type Atlas interface {
	GetImage(string) Image
}
