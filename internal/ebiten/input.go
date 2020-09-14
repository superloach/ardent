package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Input struct {
	minX, minY, maxX, maxY int
	lcx, lcy               int
	vcx, vcy               int
	capture                bool
}

func (i *Input) IsAnyKeyPressed() bool {
	for _, v := range commonToEbitenKey {
		if ebiten.IsKeyPressed(v) {
			return true
		}
	}
	return false
}

func (i *Input) IsAnyKeyJustPressed() bool {
	for _, v := range commonToEbitenKey {
		if inpututil.IsKeyJustPressed(v) {
			return true
		}
	}
	return false
}

func (i *Input) IsKeyPressed(k int) bool {
	return ebiten.IsKeyPressed(commonToEbitenKey[k])
}

func (i *Input) IsKeyJustPressed(k int) bool {
	return inpututil.IsKeyJustPressed(commonToEbitenKey[k])
}

func (i *Input) IsKeyJustReleased(k int) bool {
	return inpututil.IsKeyJustReleased(commonToEbitenKey[k])
}

func (i *Input) IsMouseButtonPressed(k int) bool {
	return ebiten.IsMouseButtonPressed(commonToEbitenMouseButton[k])
}

func (i *Input) IsMouseButtonJustPressed(k int) bool {
	return inpututil.IsMouseButtonJustPressed(commonToEbitenMouseButton[k])
}

func (i *Input) IsMouseButtonJustReleased(k int) bool {
	return inpututil.IsMouseButtonJustReleased(commonToEbitenMouseButton[k])
}

func (i *Input) CursorPosition() (int, int) {
	x, y := ebiten.CursorPosition()

	if i.minX+i.minY+i.maxX+i.maxY == 0 {
		return x, y
	}

	dx, dy := x-i.lcx, y-i.lcy
	i.lcx, i.lcy = x, y

	nx, ny := i.vcx+dx, i.vcy+dy

	switch {
	case nx < i.minX:
		i.vcx = i.minX
	case nx > i.maxX:
		i.vcx = i.maxX
	default:
		i.vcx = nx
	}

	switch {
	case ny < i.minY:
		i.vcy = i.minY
	case ny > i.maxY:
		i.vcy = i.maxY
	default:
		i.vcy = ny
	}

	return i.vcx, i.vcy
}

func (i *Input) SetCursorBounds(minX, minY, maxX, maxY int) {
	i.minX, i.minY, i.maxX, i.maxY = minX, minY, maxX, maxY
}

func (i *Input) SetCursorCapture(capture bool) {
	i.capture = capture
}
