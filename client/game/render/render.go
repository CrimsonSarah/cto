package render

import (
	"github.com/CrimsonSarah/cto/client/game/render/debug"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Renderer struct {
	World         *world.World
	CardRenderer  CardRenderer
	DebugRenderer debug.DebugRenderer
}

// Should be called after DigiGL is initialized.
func (r *Renderer) Init(world *world.World) {
	r.World = world
	r.CardRenderer.Init(world)
	r.DebugRenderer.Init(world)
}

func (r *Renderer) Configure() {
	r.CardRenderer.Configure()
	r.DebugRenderer.Configure()
}

func (r *Renderer) Clear() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
