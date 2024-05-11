package render

import "github.com/go-gl/gl/v3.3-core/gl"

type Renderer struct {
	CardRenderer CardRenderer
}

// Should be called after DigiGL is initialized
func (r *Renderer) Init() {
	r.CardRenderer.Init()
}

func (r *Renderer) Clear() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
