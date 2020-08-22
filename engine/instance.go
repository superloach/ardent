package main

type Backend byte

// Available backends
const (
	EBITEN Backend = 1 << iota
)

// Instance is an instantiated
// engine instance.
type Instance struct {
	backend Backend
}

func NewInstance(backend Backend) *Instance {
	return &Instance{backend: backend}
}
