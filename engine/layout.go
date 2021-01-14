package engine

// LayoutFunc scales the viewport, given an outside size and the virtual resolution.
//
// Basic implementations include LayoutFit, LayoutFill, and LayoutStretch.
type LayoutFunc = func(ow, oh, vw, vh int) (w, h int)

// LayoutFit is a LayoutFunc that fits the virtual resolution within the outside size.
func LayoutFit(ow, oh, vw, vh int) (w, h int) {
	return vw, vh
}

// LayoutFill is a LayoutFunc that fills the outside size with the virtual resolution.
// FIXME: doesn't work right, last I checked
func LayoutFill(ow, oh, vw, vh int) (w, h int) {
	if ow > oh {
		return vw, (vh * ow) / oh
	}

	return (vw * oh) / ow, vh
}

// LayoutStretch is a LayoutFunc that stretches to match the outside size.
func LayoutStretch(ow, oh, vw, vh int) (w, h int) {
	return ow, oh
}
