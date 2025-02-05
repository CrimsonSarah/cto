package ui

import (
	"log"

	"github.com/CrimsonSarah/cto/client/digidata"
	"github.com/CrimsonSarah/cto/client/digigl"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// This package is for creating a widget on which to create a game.
// The game itself has to be provided from somewhere else.

type Game interface {
	Init(engine.InitContext)

	// Return `false` if a redraw is NOT needed. Otherwise, return
	// `true`.
	Tick(engine.EngineFrameContext) bool
	Render(engine.EngineRenderContext)
	AddEvent(engine.Event)
}

// Ensures that the actual Game receives convenient data to work with
// and nothing more. An interface between GTK and ourselves.
type GameWrapper struct {
	Game Game

	StartTime         int64
	PreviousFrameTime int64

	EventQueue digidata.Queue[engine.Event]
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

	context := engine.InitContext{
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

		frame := engine.EngineFrameContext{
			Timei: currentFrameTime,
			Timed: currentFrameTimeS,
			Dti:   dt,
			Dtd:   dtS,
			Dtf:   float32(dtS),
			Frame: frameClock.GetFrameCounter(),
		}

		for {
			ev, ok := w.EventQueue.Dequeue()
			if !ok {
				break
			}

			w.Game.AddEvent(ev)
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
	ctx := engine.EngineRenderContext{
		GLArea:    area,
		GLContext: context,
	}

	w.Game.Render(ctx)
	return true
}

// TODO: Error handling.
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

	glArea.AddEvents(int(
		gdk.BUTTON_PRESS_MASK |
			gdk.BUTTON_RELEASE_MASK |
			gdk.KEY_PRESS_MASK |
			gdk.KEY_RELEASE_MASK |
			gdk.POINTER_MOTION_MASK,
	))

	glArea.Connect("motion-notify-event", func(area *gtk.GLArea, event *gdk.Event) {
		keyEvent := EventFromGdkPointerMotion(event)
		wrapper.EventQueue.Enqueue(keyEvent)
	})

	glArea.Connect("button-press-event", func(area *gtk.GLArea, event *gdk.Event) {
		buttonEvent := EventFromGdkButtonPress(event)
		wrapper.EventQueue.Enqueue(buttonEvent)
	})

	glArea.Connect("button-release-event", func(area *gtk.GLArea, event *gdk.Event) {
		buttonEvent := EventFromGdkButtonRelease(event)
		wrapper.EventQueue.Enqueue(buttonEvent)
	})

	glArea.SetCanFocus(true)
	glArea.Connect("key-press-event", func(area *gtk.GLArea, event *gdk.Event) {
		keyEvent := EventFromGdkKeyPress(event)
		wrapper.EventQueue.Enqueue(keyEvent)
	})

	glArea.Connect("key-release-event", func(area *gtk.GLArea, event *gdk.Event) {
		keyEvent := EventFromGdkKeyRelease(event)
		wrapper.EventQueue.Enqueue(keyEvent)
	})

	glArea.Connect("resize", func(area *gtk.GLArea, width, height int) {
		event := engine.Event{
			KindId: engine.ResizeEventKindId,
			Data: engine.ResizeEvent{
				NewWidth:  width,
				NewHeight: height,
			},
		}

		wrapper.EventQueue.Enqueue(event)
	})

	return glArea
}
