package game

import (
	"fmt"
	"log"
	"math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/objects/card"
	"github.com/CrimsonSarah/cto/client/game/render"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Game struct {
	State    GameState
	World    world.World
	Renderer render.Renderer

	Width  int
	Height int

	// Have to think harder about this.
	renderableCard1 render.RenderableCard
	renderableCard2 render.RenderableCard
}

// Goes from [0..Width] coordinates to [-1..1].
func (g *Game) normalizedWindowCoordinates(x, y float64) digimath.Vec2 {
	normalizedX := (float32(x) * 2 / float32(g.Width)) - 1
	normalizedY := -(float32(y)*2/float32(g.Height) - 1)

	return digimath.MakeVec2(normalizedX, normalizedY)
}

// Everything below is currently just for testing.

func MakeGame() Game {
	return Game{}
}

func (g *Game) Init(context ui.InitContext) {
	fmt.Println("Initializing")
	g.World = world.MakeWorld(context.Width, context.Height)
	g.Renderer.Init(&g.World)

	g.Width = context.Width
	g.Height = context.Height

	c := card.MakeCard(
		"BT5-103",
		"A Blazing Storm of Metal!",
	)

	placedCard1 := world.MakePlacedDefault(&c, &g.World)
	placedCard2 := world.MakePlacedDefault(&c, &g.World)

	// placedCard.Transform.RotateY(0.3 * math.Pi)
	placedCard1.Transform.TranslateX(-0.5)
	placedCard1.Transform.TranslateZ(-2)
	// placedCard1.Transform.Scale(1.1)

	placedCard2.Transform.TranslateX(0.5)
	placedCard2.Transform.TranslateZ(-2)

	log.Printf("Cards @ %v | %v\n",
		placedCard1.Transform.GetPosition(),
		placedCard2.Transform.GetPosition(),
	)

	log.Printf("Init renderable cards\n")
	g.renderableCard1 = g.Renderer.CardRenderer.MakeRenderableCard(&placedCard1)
	g.renderableCard2 = g.Renderer.CardRenderer.MakeRenderableCard(&placedCard2)
}

func (g *Game) Tick(f ui.FrameContext) bool {
	for event, ok := f.Events.Dequeue(); ok; event, ok = f.Events.Dequeue() {
		if eventButton, ok := event.(*gdk.EventButton); ok {
			log.Println("Event Button", eventButton)

			coords := g.normalizedWindowCoordinates(
				eventButton.X(),
				eventButton.Y(),
			)

			fmt.Printf("Point 1\n%v\n", coords)
			fmt.Printf("Intersects 1\n%v\n", world.Intersects(
				*g.renderableCard1.Placed,
				coords,
			))
			// fmt.Printf("Intersects 2\n%v\n", world.Intersects(
			// 	*g.renderableCard2.Placed,
			// 	coords,
			// ))

		} else if eventKey, ok := event.(*gdk.EventKey); ok {
			// log.Println("Event Key", eventKey, eventKey.KeyVal(), gdk.KEY_W)

			if eventKey.KeyVal() == gdk.KEY_w {
				g.renderableCard1.Transform.RotateX(0.25 * math.Pi * f.Dtf)
				g.renderableCard2.Transform.RotateX(0.25 * math.Pi * f.Dtf)
			} else if eventKey.KeyVal() == gdk.KEY_s {
				g.renderableCard1.Transform.RotateX(-0.25 * math.Pi * f.Dtf)
				g.renderableCard2.Transform.RotateX(-0.25 * math.Pi * f.Dtf)
				// fmt.Printf("Rotation %v\n", g.renderableCard1.Transform.Rotation)
			} else if eventKey.KeyVal() == gdk.KEY_a {
				g.renderableCard1.Transform.TranslateX(-0.25 * f.Dtf)
				g.renderableCard2.Transform.TranslateX(0.25 * f.Dtf)
			} else if eventKey.KeyVal() == gdk.KEY_d {
				g.renderableCard1.Transform.TranslateX(0.25 * f.Dtf)
				g.renderableCard2.Transform.TranslateX(-0.25 * f.Dtf)
			}
		}
		// log.Println("Event found!")
	}

	// g.renderableCard.Transform.RotateX(0.1 * math.Pi * f.Dtf)
	// log.Printf("Transform\n%s\n", g.renderableCard.Transform.Format())

	return true
}

func (g *Game) Render(area *gtk.GLArea, context *gdk.GLContext) {
	g.Renderer.Clear()
	g.Renderer.CardRenderer.RenderCard(&g.renderableCard1)
	g.Renderer.CardRenderer.RenderCard(&g.renderableCard2)

	g.Renderer.DebugRenderer.RenderDebug(&g.World.Debug)
}

func (g *Game) Configure(newWidth, newHeight int) {
	g.World.Configure(newWidth, newHeight)
	g.Renderer.Configure()
}
