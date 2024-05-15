package game

import (
	"log"
	"math"

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
	renderableCard1 render.RenderableCard
	renderableCard2 render.RenderableCard
}

// Everything below is currently just for testing

func MakeGame() Game {
	return Game{}
}

func (g *Game) Init(context ui.InitContext) {
	g.Renderer.Init(context.Width, context.Height)

	c := card.MakeCard(
		"BT5-103",
		"A Blazing Storm of Metal!",
	)

	placedCard1 := world.MakePlacedDefault(&c)
	placedCard2 := world.MakePlacedDefault(&c)

	// placedCard.Transform.RotateY(0.3 * math.Pi)
	placedCard1.Transform.TranslateX(-0.5)
	placedCard1.Transform.TranslateZ(-5.5)

	placedCard2.Transform.TranslateX(0.5)
	placedCard2.Transform.TranslateZ(-5.5)

	log.Printf("Init renderable cards\n")
	g.renderableCard1 = g.Renderer.CardRenderer.MakeRenderableCard(&placedCard1)
	g.renderableCard2 = g.Renderer.CardRenderer.MakeRenderableCard(&placedCard2)
}

func (g *Game) Tick(f ui.FrameContext) bool {
	for event, ok := f.Events.Dequeue(); ok; event, ok = f.Events.Dequeue() {
		if eventButton, ok := event.(*gdk.EventButton); ok {
			log.Println("Event Button", eventButton)
		} else if eventKey, ok := event.(*gdk.EventKey); ok {
			log.Println("Event Key", eventKey, eventKey.KeyVal(), gdk.KEY_W)

			if eventKey.KeyVal() == gdk.KEY_w {
				g.renderableCard1.Transform.RotateX(0.25 * math.Pi * f.Dtf)
				g.renderableCard2.Transform.RotateX(0.25 * math.Pi * f.Dtf)
			} else if eventKey.KeyVal() == gdk.KEY_s {
				g.renderableCard1.Transform.RotateX(-0.25 * math.Pi * f.Dtf)
				g.renderableCard2.Transform.RotateX(-0.25 * math.Pi * f.Dtf)
			} else if eventKey.KeyVal() == gdk.KEY_a {
				g.renderableCard1.Transform.TranslateX(-0.25 * f.Dtf)
				g.renderableCard2.Transform.TranslateX(0.25 * f.Dtf)
			} else if eventKey.KeyVal() == gdk.KEY_d {
				g.renderableCard1.Transform.TranslateX(0.25 * f.Dtf)
				g.renderableCard2.Transform.TranslateX(-0.25 * f.Dtf)
			}
		}
		log.Println("Event found!")
	}

	// g.renderableCard.Transform.RotateX(0.1 * math.Pi * f.Dtf)
	// log.Printf("Transform\n%s\n", g.renderableCard.Transform.Format())

	return true
}

func (g *Game) Render(area *gtk.GLArea, context *gdk.GLContext) {
	g.Renderer.Clear()
	g.Renderer.CardRenderer.RenderCard(&g.renderableCard1)
	g.Renderer.CardRenderer.RenderCard(&g.renderableCard2)
}

func (g *Game) Configure(newWidth, newHeight int) {
	g.Renderer.Configure(newWidth, newHeight)
}
