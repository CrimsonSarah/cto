package game

import (
	"log"
	"math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/card"
	"github.com/CrimsonSarah/cto/client/game/render"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Game struct {
	State    GameState
	Renderer render.Renderer

	// Have to think harder about this
	renderableCard render.RenderableCard
}

// Everything below is currently just for testing

func MakeGame() Game {
	return Game{}
}

func (g *Game) Init() {
	g.Renderer.Init()

	c := card.MakeCard(
		"BT5-103",
		"A Blazing Storm of Metal!",
	)

	placed_card := world.MakePlacedDefault(&c)

	log.Printf("Init renderable card\n")
	g.renderableCard = g.Renderer.CardRenderer.MakeRenderableCard(&placed_card)
}

func (g *Game) Tick(time ui.GameTime) bool {
	g.renderableCard.Transform = g.renderableCard.Transform.Mul(
		digimath.Matrix44RotateZ(math.Pi * time.Dtf),
	)

	return true
}

func (g *Game) Render(area *gtk.GLArea, context *gdk.GLContext) {
	g.Renderer.Clear()
	g.Renderer.CardRenderer.RenderCard(&g.renderableCard)
}
