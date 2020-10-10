package engine

type Input interface {
	// Keyboard
	IsAnyKeyPressed() bool
	IsAnyKeyJustPressed() bool
	IsKeyPressed(int) bool
	IsKeyJustPressed(int) bool
	IsKeyJustReleased(int) bool

	// Mouse
	IsMouseButtonPressed(int) bool
	IsMouseButtonJustPressed(int) bool
	IsMouseButtonJustReleased(int) bool
	CursorPosition() (int, int)
	SetCursorBounds(int, int, int, int)
	SetCursorMode(CursorMode)
}

// CursorMode indication
type CursorMode byte

// Cursor modes
const (
	CURSOR_VISIBLE CursorMode = 1 << iota
	CURSOR_HIDDEN
	CURSOR_CAPTURED
)
