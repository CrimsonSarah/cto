package ui

import (
	"log"

	"github.com/CrimsonSarah/cto/client/digigl"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// This package is for creating a widget on which to create a game.
// The game itself has to be provided from somewhere else.

type GameTime struct {
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

type Game interface {
	Init()

	// Return `false` if a redraw is NOT needed. Otherwise, return
	// `true`.
	Tick(GameTime) bool

	Render(area *gtk.GLArea, context *gdk.GLContext)
}

// Ensures that the actual Game receives convenient data to work with
// and nothing more. An interface between GTK and ourselves.
type GameWrapper struct {
	Game Game

	StartTime         int64
	PreviousFrameTime int64
}

type InitCallback func()

func (w *GameWrapper) Init(area *gtk.GLArea) {
	area.MakeCurrent()

	if err := area.GetError(); err != nil {
		log.Fatalln("Could not create make GL area current.", err)
		return
	}

	w.StartTime = area.GetFrameClock().GetFrameTime()
	w.PreviousFrameTime = 0

	digigl.DigiGLInit()
	w.Game.Init()
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

		gameTime := GameTime{
			Timei: currentFrameTime,
			Timed: currentFrameTimeS,
			Dti:   dt,
			Dtd:   dtS,
			Dtf:   float32(dtS),
			Frame: frameClock.GetFrameCounter(),
		}

		shouldDraw := w.Game.Tick(gameTime)

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

	return glArea
}
