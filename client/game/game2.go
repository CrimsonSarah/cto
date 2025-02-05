package game

import (
	"log"
	"math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/camera"
	"github.com/CrimsonSarah/cto/client/game/card"
	"github.com/CrimsonSarah/cto/client/game/generic/collider"
	"github.com/CrimsonSarah/cto/client/game/generic/sprite"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Game struct {
	World world.World

	Width  int
	Height int

	Engine *engine.Engine[global.GameGlobals]
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

func (g *Game) Init(context engine.InitContext) {
	log.Println("Initializing")
	g.World = world.MakeWorld(context.Width, context.Height)
	g.Engine = engine.NewEngine[global.GameGlobals](&g.World)

	g.Width = context.Width
	g.Height = context.Height

	e := g.Engine

	collider.AddCollisionSystems(e)
	transform.AddTransformSystems(e)
	sprite.AddSpriteSystems(e)
	camera.AddCameraSystems(e)

	camera, err := camera.AddCamera(e, g.Width, g.Height)
	e.GetGlobals().Camera = camera.Id

	card1, err := card.AddCard(e)

	if err != nil {
		log.Printf("Error creating card: %s\n", err.Error())
	}

	e.UpdateComponent(
		card1,
		&transform.TransformComponentKind,
		func(c engine.Component) engine.Component {
			t := c.Data.(transform.Transform)
			t.TranslateX(-0.5)
			t.TranslateZ(-2)

			log.Printf("Card 1 @ %v\n",
				t.GetPosition(),
			)

			c.Data = t
			return c
		},
	)

	card2, err := card.AddCard(e)

	if err != nil {
		log.Printf("Error creating card: %s\n", err.Error())
	}

	e.UpdateComponent(
		card2,
		&transform.TransformComponentKind,
		func(c engine.Component) engine.Component {
			t := c.Data.(transform.Transform)
			t.TranslateX(0.5)
			t.TranslateZ(-2)
			t.RotateY(math.Pi / 2)

			log.Printf("Card 2 @ %v\n",
				t.GetPosition(),
			)

			c.Data = t
			return c
		},
	)

	g.World.AddClipPoint(
		digimath.MakeVec2(0, 0),
		digimath.MakeVec3(0.5, 0.5, 0.5),
	)
}

func (g *Game) AddEvent(ev engine.Event) {
	g.Engine.AddEvent(ev)
}

func (g *Game) Tick(ctx engine.EngineFrameContext) bool {
	g.Engine.Tick(ctx)
	return true
}

func (g *Game) Render(ctx engine.EngineRenderContext) {
	// TODO: Extract.
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	g.Engine.Render(ctx)
}

// func (g *Game) Configure(newWidth, newHeight int) {
// 	g.World.Configure(newWidth, newHeight)
// 	g.Renderer.Configure()
// }
