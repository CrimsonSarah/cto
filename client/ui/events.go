package ui

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/gotk3/gotk3/gdk"
)

// FYI apparently the Type() of the GDK events goes away when their
// corresponding event handler exits.

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

var PointerMotionEventKindId = engine.MakeEventKindId()

type PointerMotionEvent struct {
	pointerEvent
}

var PointerButtonDownEventKindId = engine.MakeEventKindId()
var PointerButtonUpEventKindId = engine.MakeEventKindId()

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

var KeyDownEventKindId = engine.MakeEventKindId()
var KeyUpEventKindId = engine.MakeEventKindId()

type KeyDownEvent struct {
	keyEvent
}

type KeyUpEvent struct {
	keyEvent
}

// Conversions.

var buttonMap = map[gdk.Button]PointerButton{
	1: PointerButton1,
	2: PointerButton2,
	3: PointerButton3,
	4: PointerButton4,
	5: PointerButton5,
}

func EventFromGdkPointerMotion(ev *gdk.Event) engine.Event {
	event := gdk.EventMotionNewFromEvent(ev)
	x, y := event.MotionVal()

	return engine.Event{
		KindId: PointerMotionEventKindId,
		Data: PointerMotionEvent{
			pointerEvent: pointerEvent{
				X: float32(x),
				Y: float32(y),
			},
		},
	}
}

func EventFromGdkButtonPress(ev *gdk.Event) engine.Event {
	event := gdk.EventButtonNewFromEvent(ev)

	return engine.Event{
		KindId: PointerButtonDownEventKindId,
		Data: PointerButtonDownEvent{
			pointerEvent: pointerEvent{
				X: float32(event.X()),
				Y: float32(event.Y()),
			},
			Button: buttonMap[event.Button()],
		},
	}
}

func EventFromGdkButtonRelease(ev *gdk.Event) engine.Event {
	event := gdk.EventButtonNewFromEvent(ev)

	return engine.Event{
		KindId: PointerButtonUpEventKindId,
		Data: PointerButtonUpEvent{
			pointerEvent: pointerEvent{
				X: float32(event.X()),
				Y: float32(event.Y()),
			},
			Button: buttonMap[event.Button()],
		},
	}
}

func EventFromGdkKeyPress(ev *gdk.Event) engine.Event {
	event := gdk.EventKeyNewFromEvent(ev)

	return engine.Event{
		KindId: KeyDownEventKindId,
		Data: KeyDownEvent{
			keyEvent: keyEvent{
				Key:       gdk.KeyvalToUpper(event.KeyVal()),
				Modifiers: event.State(),
			},
		},
	}
}

func EventFromGdkKeyRelease(ev *gdk.Event) engine.Event {
	event := gdk.EventKeyNewFromEvent(ev)

	return engine.Event{
		KindId: KeyUpEventKindId,
		Data: KeyUpEvent{
			keyEvent: keyEvent{
				Key:       gdk.KeyvalToUpper(event.KeyVal()),
				Modifiers: event.State(),
			},
		},
	}
}
