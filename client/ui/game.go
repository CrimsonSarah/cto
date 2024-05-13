package ui

import (
	"log"

	"github.com/CrimsonSarah/cto/client/digidata"
	"github.com/CrimsonSarah/cto/client/digigl"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// This package is for creating a widget on which to create a game.
// The game itself has to be provided from somewhere else.

type FrameContext struct {
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

	// See https://docs.gtk.org/gtk3/input-handling.html
	// Use type assertions to determine the event type. They will be
	// pointers!
	// Feel compelled to remove events once they are handled.
	Events *digidata.Queue[interface{}]
}

type InitContext struct {
	Width  int
	Height int
}

type Game interface {
	Init(InitContext)

	// Return `false` if a redraw is NOT needed. Otherwise, return
	// `true`.
	Tick(FrameContext) bool

	Render(area *gtk.GLArea, context *gdk.GLContext)
}

// Ensures that the actual Game receives convenient data to work with
// and nothing more. An interface between GTK and ourselves.
type GameWrapper struct {
	Game Game

	StartTime         int64
	PreviousFrameTime int64

	// Not sure if there is the need to do this manually. Stores events
	// to make sure they are processed deterministically every frame.
	// Use type assertions to determine the event type.
	EventQueue digidata.Queue[interface{}]
}

func (w *GameWrapper) Init(area *gtk.GLArea) {
	area.MakeCurrent()

	if err := area.GetError(); err != nil {
		log.Fatalln("Could not create make GL area current.", err)
		return
	}

	w.StartTime = area.GetFrameClock().GetFrameTime()
	w.PreviousFrameTime = 0

	digigl.DigiGLInit()

	context := InitContext{
		Width:  area.GetAllocatedWidth(),
		Height: area.GetAllocatedHeight(),
	}

	w.Game.Init(context)
}

func (w *GameWrapper) Tick(widget *gtk.Widget, frameClock *gdk.FrameClock) bool {
	generic, _ := widget.Cast()

	if area, ok := generic.(*gtk.GLArea); ok {
		monotonicFrameTime := frameClock.GetFrameTime()

		currentFrameTime := monotonicFrameTime - w.StartTime
		currentFrameTimeS := float64(currentFrameTime) / 1e6

		dt := currentFrameTime - w.PreviousFrameTime
		dtS := float64(dt) / 1e6

		w.PreviousFrameTime = currentFrameTime

		frame := FrameContext{
			Timei:  currentFrameTime,
			Timed:  currentFrameTimeS,
			Dti:    dt,
			Dtd:    dtS,
			Dtf:    float32(dtS),
			Frame:  frameClock.GetFrameCounter(),
			Events: &w.EventQueue,
		}

		shouldDraw := w.Game.Tick(frame)

		if shouldDraw {
			area.QueueDraw()
		}

		return true
	} else {
		log.Fatalln("Widget passed to callback is not a GLArea..?")
		return false
	}
}

func (w *GameWrapper) Render(area *gtk.GLArea, context *gdk.GLContext) bool {
	w.Game.Render(area, context)
	return true
}

// TODO: Error handling
func GameWidgetNew(game Game) *gtk.GLArea {
	glArea, err := gtk.GLAreaNew()

	if err != nil {
		log.Fatalln("Could not create GL area.", err)
		return nil
	}

	wrapper := GameWrapper{
		Game: game,
	}

	glArea.Connect("realize", wrapper.Init)
	glArea.Connect("render", wrapper.Render)
	glArea.AddTickCallback(wrapper.Tick)

	glArea.AddEvents(int(gdk.BUTTON_PRESS_MASK | gdk.KEY_PRESS_MASK))

	glArea.Connect("button-press-event", func(area *gtk.GLArea, event *gdk.Event) {
		buttonEvent := gdk.EventButtonNewFromEvent(event)
		wrapper.EventQueue.Enqueue(buttonEvent)
	})

	glArea.SetCanFocus(true)
	glArea.Connect("key-press-event", func(area *gtk.GLArea, event *gdk.Event) {
		keyEvent := gdk.EventKeyNewFromEvent(event)
		wrapper.EventQueue.Enqueue(keyEvent)
	})

	return glArea
}
