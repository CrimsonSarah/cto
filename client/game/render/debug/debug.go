package debug

import (
	"github.com/CrimsonSarah/cto/client/game/world"
)

type DebugRenderer struct {
	DebugLineRenderer  DebugLineRenderer
	DebugPointRenderer DebugPointRenderer
}

// Should be called after DigiGL is initialized
func (r *DebugRenderer) Init(world *world.World) {
	r.DebugLineRenderer.Init(world)
	r.DebugPointRenderer.Init(world)
}

func (r *DebugRenderer) Configure() {
	r.DebugLineRenderer.Configure()
	r.DebugPointRenderer.Configure()
}

// Encapsulating the WorldDebug is not needed in this case.
func (r *DebugRenderer) RenderDebug(debug *world.WorldDebug) {
	for _, lineGroup := range debug.LineGroups {
		renderableLines := r.DebugLineRenderer.MakeRenderableDebugLines(
			lineGroup,
		)

		r.DebugLineRenderer.RenderDebugLine(&renderableLines)
	}

	for _, point := range debug.Points {
		renderablePoint := r.DebugPointRenderer.MakeRenderableDebugPoint(
			point,
		)

		r.DebugPointRenderer.RenderDebugPoint(&renderablePoint)
	}
}
