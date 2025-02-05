package camera

import (
	"math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
)

type CameraExploreSystemImpl struct{}

var CameraExploreSystem = engine.MakeSystem[global.GameGlobals](
	"CameraExplore",
	CameraExploreSystemImpl{},
)

func (s CameraExploreSystemImpl) EventKinds() []engine.EventKindId {
	return []engine.EventKindId{
		engine.FrameEventKindId,
	}
}

func (s CameraExploreSystemImpl) NodeKind() *engine.NodeKind {
	return &CameraExploreNodeKind
}

func (s CameraExploreSystemImpl) Init(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
}

func (s CameraExploreSystemImpl) Handle(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	transformC :=
		node.GetComponent(&transform.TransformComponentKind)
	transform := transformC.Data.(transform.Transform)

	inputC :=
		node.GetComponent(&CameraInputComponentKind)
	input := inputC.Data.(CameraInput)

	if input.InputF1 {
		CameraToDefaultState(ctx, node.Entity)
		return
	}

	// Rotation.

	// Movement along the X axis corresponds to rotation around the
	// Y axis and movement along the Y axis corresponds to rotation
	// around the X axis.
	ry := -input.PointerPosition.X() * 2 * math.Pi
	rx := input.PointerPosition.Y() * 2 * math.Pi

	transform.Rotation = digimath.MakeVec3(rx, ry, 0)

	// Translation.

	speed := float32(2)

	forward := transform.GetForwardFlat()
	left := digimath.Matrix33RotateY(math.Pi / 2).MulV(forward)
	up := digimath.MakeVec3(0, 1, 0)
	direction := digimath.Vec3Zero

	if input.InputForward {
		direction = direction.Add(forward)
	}
	if input.InputBackward {
		direction = direction.Add(forward.Scale(-1))
	}
	if input.InputLeft {
		direction = direction.Add(left)
	}
	if input.InputRight {
		direction = direction.Add(left.Scale(-1))
	}
	if input.InputUp {
		direction = direction.Add(up)
	}
	if input.InputDown {
		direction = direction.Add(up.Scale(-1))
	}

	frame := event.(engine.FrameEvent)
	displacement := direction.Scale(speed * frame.Dtf)

	transform.Translate(displacement)

	ctx.OverwriteComponent(node.Entity, transformC.Update(transform))
}
