package engine

type BasicRenderer interface {
	AddImage(...Image)
	Draw()
}
