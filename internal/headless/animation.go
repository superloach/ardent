package headless

type Animation struct {
	Image
}

func (a Animation) SetState(state string) {}

func (a Animation) Play() {}

func (a Animation) Pause() {}

func (a Animation) Reset() {}
