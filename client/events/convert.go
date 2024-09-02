package events

import (
	"github.com/gotk3/gotk3/gdk"
)

var buttonMap = map[gdk.Button]PointerButton{
	1: PointerButton1,
	2: PointerButton2,
	3: PointerButton3,
	4: PointerButton4,
	5: PointerButton5,
}

func FromGdkPointerMotion(ev *gdk.Event) PointerMotionEvent {
	event := gdk.EventMotionNewFromEvent(ev)
	x, y := event.MotionVal()

	return PointerMotionEvent{
		pointerEvent: pointerEvent{
			X: float32(x),
			Y: float32(y),
		},
	}
}

func FromGdkButtonPress(ev *gdk.Event) PointerButtonDownEvent {
	event := gdk.EventButtonNewFromEvent(ev)

	return PointerButtonDownEvent{
		pointerEvent: pointerEvent{
			X: float32(event.X()),
			Y: float32(event.Y()),
		},
		Button: buttonMap[event.Button()],
	}
}

func FromGdkButtonRelease(ev *gdk.Event) PointerButtonUpEvent {
	event := gdk.EventButtonNewFromEvent(ev)

	return PointerButtonUpEvent{
		pointerEvent: pointerEvent{
			X: float32(event.X()),
			Y: float32(event.Y()),
		},
		Button: buttonMap[event.Button()],
	}
}

func FromGdkKeyPress(ev *gdk.Event) KeyDownEvent {
	event := gdk.EventKeyNewFromEvent(ev)

	return KeyDownEvent{
		keyEvent: keyEvent{
			Key:       gdk.KeyvalToUpper(event.KeyVal()),
			Modifiers: event.State(),
		},
	}
}

func FromGdkKeyRelease(ev *gdk.Event) KeyUpEvent {
	event := gdk.EventKeyNewFromEvent(ev)

	return KeyUpEvent{
		keyEvent: keyEvent{
			Key:       gdk.KeyvalToUpper(event.KeyVal()),
			Modifiers: event.State(),
		},
	}
}
