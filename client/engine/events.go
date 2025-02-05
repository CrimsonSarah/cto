package engine

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type EventKindId int

var eventKindIdCounter = 0

func MakeEventKindId() EventKindId {
	id := eventKindIdCounter
	eventKindIdCounter += 1

	return EventKindId(id)
}

type Event struct {
	KindId EventKindId
	Data any
}

// Default events (feel free to create more).

// Frame. Should happen at a fixed rate.

var FrameEventKindId = MakeEventKindId()

type FrameEvent struct {
	// Time in microsseconds from start.
	Timei int64
	// Time in seconds from start (64 bits).
	Timed float64
	// Time in microsseconds from last frame.
	Dti int64
	// Time in seconds from last frame (64 bits).
	Dtd float64
	// Time in seconds from last frame (32 bits).
	Dtf float32
	// What frame this is.
	Frame int64
}

// Render. Should happen every time the outside system asks us to draw.

var RenderEventKindId = MakeEventKindId()

type RenderEvent struct {
	GLArea    *gtk.GLArea
	GLContext *gdk.GLContext
}

// Resize. Happens on window resize.

var ResizeEventKindId = MakeEventKindId()

type ResizeEvent struct {
	NewWidth int
	NewHeight int
}

// EntityCreated. Is enqueued every time an Entity is added to the engine.

var EntityCreatedEventKindId = MakeEventKindId()

type EntityCreatedEvent struct {
	EntityId EntityId
}
