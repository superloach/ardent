package engine

type Entity interface {
	Tick()

	SetCollider(Collider)

	Position() Vec2

	AddImage(...Image)
	Images() []Image

	Class() string

	Dispose()
	IsDisposed() bool
}

type CoreEntity struct {
	Vec2
	prevPos Vec2

	images []Image

	collider Collider
	disposed bool
}

func (e *CoreEntity) Tick() {
	if e.collider != nil {
		e.Vec2 = e.collider.Resolve(e.prevPos, e.Vec2)
	}
	e.prevPos = e.Vec2

	for _, img := range e.images {
		img.Translate(e.X, e.Y)
	}
}

func (e *CoreEntity) SetCollider(collider Collider) {
	e.collider = collider
}

func (e *CoreEntity) Position() Vec2 {
	return e.Vec2
}

func (e *CoreEntity) AddImage(image ...Image) {
	for _, img := range image {
		img.Translate(e.X, e.Y)
	}
	e.images = append(e.images, image...)
}

func (e *CoreEntity) Images() []Image {
	return e.images
}

func (e *CoreEntity) Dispose() {
	e.disposed = true
	for _, img := range e.images {
		img.Dispose()
	}
}

func (e *CoreEntity) IsDisposed() bool {
	return e.disposed
}
