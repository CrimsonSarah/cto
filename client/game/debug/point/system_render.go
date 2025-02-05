package point

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/camera"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type DebugPointRenderSystemImpl struct{}

var DebugPointRenderSystem = engine.MakeSystem[global.GameGlobals](
	"DebugPointRender",
	DebugPointRenderSystemImpl{},
)

func (s DebugPointRenderSystemImpl) EventKinds() []engine.EventKindId {
	return []engine.EventKindId{engine.RenderEventKindId}
}

func (s DebugPointRenderSystemImpl) NodeKind() *engine.NodeKind {
	return &DebugPointRenderNodeKind
}

func (s DebugPointRenderSystemImpl) Init(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	render := node.GetData(&DebugPointComponentKind).(DebugPoint)
	render.Bind()

	camera := camera.GetCameraCamera(ctx)

	gl.UniformMatrix4fv(
		render.ProjectionUniformLocation,
		1,
		false,
		&camera.Projection[0],
	)
}

func (s DebugPointRenderSystemImpl) Handle(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	transform :=
		node.GetData(&transform.TransformComponentKind).(transform.Transform)
	render :=
		node.GetData(&DebugPointComponentKind).(DebugPoint)

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
