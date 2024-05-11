// Reexport
package ui

import (
	"github.com/CrimsonSarah/cto/client/ui/game"
	"github.com/gotk3/gotk3/gtk"
)

func GameWidgetNew(g game.Game) *gtk.GLArea {
	return game.GameWidgetNew(g)
}
