package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "com.github.crimsonsarah.cto"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	application.Connect("activate", func() { onActivate(application) })
	os.Exit(application.Run(os.Args))
}

func onActivate(application *gtk.Application) {
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

	l, err := gtk.LabelNew("Hello!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	appWindow.Add(l)
	appWindow.ShowAll()
}
