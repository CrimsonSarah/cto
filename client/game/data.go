package game

import (
	"github.com/CrimsonSarah/cto/client/game/render"
)

// Source of truth.
type GameData struct {
	// What cards are being seen. Used to keep track of loaded
	// textures, and possibly other things.
	Cards      []render.RenderableCard
	CardCounts map[string]uint8 // Keys are codes, not IDs

	Hand []string // These are IDs
}
