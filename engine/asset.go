package engine

import "io"

type Asset interface {
	ToImage() (Image, error)

	io.ReadWriter
}
