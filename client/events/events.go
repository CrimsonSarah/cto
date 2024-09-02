package events

// Using handmade events because it seems the Type() of the GDK ones
// goes away when their corresponding event handler exits.

type PointerButton uint8

const (
	PointerButton1 = PointerButton(iota)
	PointerButton2
	PointerButton3
	PointerButton4
	PointerButton5
)

type pointerEvent struct {
	X float32
	Y float32
}

type PointerMotionEvent struct {
	pointerEvent
}

type PointerButtonDownEvent struct {
	pointerEvent
	Button PointerButton
}

type PointerButtonUpEvent struct {
	pointerEvent
	Button PointerButton
}

type keyEvent struct {
	// Key values are found in `gdk`.
	// Currently, every key is converted to upper case. So check for
	// KEY_D and not KEY_d.
	Key       uint
	Modifiers uint // Masks are found in `gdk`.
}

type KeyDownEvent struct {
	keyEvent
}

type KeyUpEvent struct {
	keyEvent
}
