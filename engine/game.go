package engine

// Backend flag type
type Backend byte

// Available backends
const (
	EBITEN Backend = 1 << iota
)

// Game is an instantiated
// engine instance.
type Game struct {
	backend Backend
}

func NewGame(backend Backend) *Instance {
	return &Game{backend: backend}
}
