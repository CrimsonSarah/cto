package sprite

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/camera"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type SpriteRenderSystemImpl struct{}

var SpriteRenderSystem = engine.MakeSystem[global.GameGlobals](
	"SpriteRender",
	SpriteRenderSystemImpl{},
)

func (s SpriteRenderSystemImpl) EventKinds() []engine.EventKindId {
	return []engine.EventKindId{engine.RenderEventKindId}
}

func (s SpriteRenderSystemImpl) NodeKind() *engine.NodeKind {
	return &SpriteRenderNodeKind
}

func (s SpriteRenderSystemImpl) Init(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	render := node.GetData(&SpriteComponentKind).(Sprite)
	render.Bind()

	camera := camera.GetCameraCamera(ctx)

	gl.UniformMatrix4fv(
		render.ProjectionUniformLocation,
		1,
		false,
		&camera.Projection[0],
	)
}

func (s SpriteRenderSystemImpl) Handle(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	transform :=
		node.GetData(&transform.TransformComponentKind).(transform.Transform)
	render :=
		node.GetData(&SpriteComponentKind).(Sprite)

	render.Bind()

	transformM := transform.ToMatrix()
	cameraM := camera.GetCameraViewMatrix(ctx)

	gl.UniformMatrix4fv(
		render.CameraUniformLocation,
		1,
		false,
		&cameraM[0],
	)

	gl.UniformMatrix4fv(
		render.TransformUniformLocation,
		1,
		false,
		&transformM[0],
	)

	gl.DrawElements(
		gl.TRIANGLES,
		render.Mesh.IndexCount(),
		gl.UNSIGNED_INT,
		nil,
	)
}
