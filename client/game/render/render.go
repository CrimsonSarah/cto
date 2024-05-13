package render

import (
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Renderer struct {
	CardRenderer CardRenderer
}

// Should be called after DigiGL is initialized
func (r *Renderer) Init(width, height int) {
	projection := world.GetProjection(
		float32(width),
		float32(height),
	)

	r.CardRenderer.Init(&projection)
}

func (r *Renderer) Configure(width, height int) {
	projection := world.GetProjection(
		float32(width),
		float32(height),
	)

	r.CardRenderer.Configure(&projection)
}

func (r *Renderer) Clear() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
