package game

import (
	"log"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func openGlDebug(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message string,
	userParam unsafe.Pointer) {
	log.Printf("[OpenGL Debug Message] source: %d, type: %d, id: %d, severity: %d, message: %s\n",
		source, gltype, id, severity, message)
}

type InitCallback func(area *gtk.GLArea)

func wrapOnRealize(callback InitCallback) InitCallback {
	return func(area *gtk.GLArea) {
		area.MakeCurrent()

		if err := area.GetError(); err != nil {
			log.Fatalln("Could not create make GL area current.", err)
			return
		}

		if err := gl.Init(); err != nil {
			log.Fatalln("Could not initialize OpenGL", err)
			return
		}

		gl.Enable(gl.DEBUG_OUTPUT)
		gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)

		gl.DebugMessageCallback(openGlDebug, nil)
		gl.DebugMessageControl(gl.DONT_CARE, gl.DONT_CARE, gl.DONT_CARE, 0, nil, true)

		callback(area)
	}
}

type RenderCallback func(area *gtk.GLArea, context *gdk.GLContext) bool

func wrapOnRender(callback RenderCallback) RenderCallback {
	return func(area *gtk.GLArea, context *gdk.GLContext) bool {
		// FIXME: For some reason I don't seem to be able to move these
		// to a function, or anywhere else really, without causing a
		// segfault... (?)
		gl.ClearColor(0, 0, 0, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		callback(area, context)

		return true
	}
}

type TickCallback func(area *gtk.GLArea, frameClock *gdk.FrameClock) bool

func wrapOnTick(callback TickCallback) gtk.TickCallback {
	return func(widget *gtk.Widget, frameClock *gdk.FrameClock) bool {
		generic, _ := widget.Cast()

		if area, ok := generic.(*gtk.GLArea); ok {
			return callback(area, frameClock)
		} else {
			log.Fatalln("Widget passed to callback is not a GLArea..?")
			return false
		}
	}
}

func GameNew(init InitCallback, tick TickCallback, render RenderCallback) *gtk.GLArea {
	glArea, err := gtk.GLAreaNew()

	if err != nil {
		log.Fatalln("Could not create GL area.", err)
	}

	glArea.SetSizeRequest(100, 100)
	glArea.Connect("realize", wrapOnRealize(init))
	glArea.Connect("render", wrapOnRender(render))

	glArea.AddTickCallback(wrapOnTick(tick))
	return glArea
}
