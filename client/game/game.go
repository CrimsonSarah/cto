package game

import (
	"fmt"
	"log"

	"github.com/CrimsonSarah/cto/client/digidata"
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/events"
	"github.com/CrimsonSarah/cto/client/game/objects/card"
	"github.com/CrimsonSarah/cto/client/game/render"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Game struct {
	Data     GameData
	World    world.World
	Renderer render.Renderer

	State GameState

	Width  int
	Height int

	// Have to think harder about this.
	renderableCard1 render.RenderableCard
	renderableCard2 render.RenderableCard
}

// Goes from [0..Width] coordinates to [-1..1].
func (g *Game) normalizedWindowCoordinates(x, y float32) digimath.Vec2 {
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

	state := GameDefaultState{}
	g.State = &state
	g.State.Init(g)

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
	// Unhandled root events will go here for the current state to
	// handle.
	stateEvents := digidata.MakeQueue[any]()

	for ev, ok := f.Events.Dequeue(); ok; ev, ok = f.Events.Dequeue() {
		if event, ok := ev.(events.PointerButtonDownEvent); ok {
			coords := g.normalizedWindowCoordinates(
				event.X,
				event.Y,
			)

			fmt.Printf("Intersects 1\n%v\n", world.Intersects(
				*g.renderableCard1.Placed,
				coords,
			))

			stateEvents.Enqueue(ev)
		} else {
			stateEvents.Enqueue(ev)
		}
	}

	f.Events = &stateEvents
	g.State.Tick(f)

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
