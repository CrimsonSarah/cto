package main

import (
	"log"
	"math"
	"os"
	"runtime"

	"github.com/CrimsonSarah/cto/client/card"
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/CrimsonSarah/cto/client/ui/render"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	log.Println("Started!")

	const appID = "com.github.crimsonsarah.cto"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	application.Connect("activate", func() { onActivate(application) })
	os.Exit(application.Run(os.Args))
}

func onActivate(application *gtk.Application) {
	log.Println("Activated!")

	appWindow, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal("Could not create application window.", err)
	}

	appWindow.SetTitle("CTO")
	appWindow.SetDefaultSize(640, 480)

	css, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Unable to create CSS provider:", err)
	}

	// https://github.com/gotk3/gotk3/pull/926
	css.LoadFromData(`
		window { background-color: black; }
		label { color: yellow; }
	`)

	gtk.AddProviderForScreen(appWindow.GetScreen(), css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// TEST
	card := card.MakePlacedCard(
		"BT5-103",
		"A Blazing Storm of Metal!",
	)

	var renderableCard render.RenderableCard
	var previousFrameTime int64

	onInit := func(area *gtk.GLArea) {
		render.RenderInit()
		renderableCard = render.MakeRenderableCard(&card)

		previousFrameTime = area.GetFrameClock().GetFrameTime()
	}

	onUpdate := func(area *gtk.GLArea, frameClock *gdk.FrameClock) bool {
		currentFrameTime := frameClock.GetFrameTime()
		dt := currentFrameTime - previousFrameTime
		dtS := float32(dt) / 1e6

		previousFrameTime = currentFrameTime

		renderableCard.Transform = renderableCard.Transform.Mul(
			digimath.Matrix44RotateZ(math.Pi * dtS),
		)

		area.QueueDraw()
		return true
	}

	onRender := func(area *gtk.GLArea, context *gdk.GLContext) bool {
		log.Println("Rendering!")
		render.RenderCard(&renderableCard)

		return true
	}

	game := ui.GameNew(onInit, onUpdate, onRender)

	appWindow.Add(game)
	appWindow.ShowAll()
}
