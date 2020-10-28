package engine

type Animation interface {
	SetState(string)

	Play()
	Pause()
	Reset()

	Image
}
