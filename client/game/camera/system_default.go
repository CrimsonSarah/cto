package camera

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
)

type CameraDefaultSystemImpl struct{}

var CameraDefaultSystem = engine.MakeSystem[global.GameGlobals](
	"CameraDefault",
	CameraDefaultSystemImpl{},
)

func (s CameraDefaultSystemImpl) EventKinds() []engine.EventKindId {
	return []engine.EventKindId{
		engine.FrameEventKindId,
	}
}

func (s CameraDefaultSystemImpl) NodeKind() *engine.NodeKind {
	return &CameraDefaultNodeKind
}

func (s CameraDefaultSystemImpl) Init(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	transformC :=
		node.GetComponent(&transform.TransformComponentKind)
	transform := transformC.Data.(transform.Transform)

	transform.ScaleFactor = 1
	transform.Rotation = digimath.Vec3Zero
	transform.Position = digimath.Vec3Zero

	ctx.OverwriteComponent(node.Entity, transformC.Update(transform))
}

func (s CameraDefaultSystemImpl) Handle(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	inputC :=
		node.GetComponent(&CameraInputComponentKind)
	input := inputC.Data.(CameraInput)

	if input.InputF2 {
		CameraToExploreState(ctx, node.Entity)
		return
	}
}
