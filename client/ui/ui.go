// Reexport
package ui

import (
	"github.com/CrimsonSarah/cto/client/ui/game"
	"github.com/gotk3/gotk3/gtk"
)

func GameNew(
	init game.InitCallback,
	tick game.TickCallback,
	render game.RenderCallback,
) *gtk.GLArea {
	return game.GameNew(init, tick, render)
}
