package headless

type Input struct{}

func (i Input) IsAnyKeyPressed() bool {
	return false
}

func (i Input) IsAnyKeyJustPressed() bool {
	return false
}

func (i Input) IsKeyPressed(k int) bool {
	return false
}

func (i Input) IsKeyJustPressed(k int) bool {
	return false
}

func (i Input) IsKeyJustReleased(k int) bool {
	return false
}

func (i Input) IsMouseButtonPressed(k int) bool {
	return false
}

func (i Input) IsMouseButtonJustPressed(k int) bool {
	return false
}

func (i Input) IsMouseButtonJustReleased(k int) bool {
	return false
}

func (i Input) CursorPosition() (int, int) {
	return 0, 0
}

func (i Input) SetCursorBounds(minX, minY, maxX, maxY int) {}

func (i Input) SetCursorCapture(capture bool) {}
