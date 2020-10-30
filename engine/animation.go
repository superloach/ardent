package engine

type Animation interface {
	SetState(string)

	SetTickCount(int)
	Play()
	Pause()
	Reset()

	Image
}
