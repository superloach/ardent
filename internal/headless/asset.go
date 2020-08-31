package headless

import "github.com/split-cube-studios/ardent/engine"

type Asset struct{}

func (a Asset) ToImage() (engine.Image, error) {
	return new(Image), nil
}
func (a Asset) Read(p []byte) (int, error) {
	return 0, nil
}

func (a Asset) Write(p []byte) (int, error) {
	return len(p), nil
}
