package camera

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/gotk3/gotk3/gdk"
)

type CameraInputSystemImpl struct{}

var CameraInputSystem = engine.MakeSystem[global.GameGlobals](
	"CameraInput",
	CameraInputSystemImpl{},
)

func (s CameraInputSystemImpl) EventKinds() []engine.EventKindId {
	return []engine.EventKindId{
		ui.PointerMotionEventKindId,
		ui.PointerButtonDownEventKindId,
		ui.KeyDownEventKindId,
		ui.KeyUpEventKindId,
	}
}

func (s CameraInputSystemImpl) NodeKind() *engine.NodeKind {
	return &CameraInputNodeKind
}

func (s CameraInputSystemImpl) Init(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
}

func (s CameraInputSystemImpl) Handle(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	inputC :=
		node.GetComponent(&CameraInputComponentKind)

	input := inputC.Data.(CameraInput)

	cameraC :=
		node.GetComponent(&CameraCameraComponentKind)

	camera := cameraC.Data.(CameraCamera)

	if ev, ok := event.(ui.KeyDownEvent); ok {
		if ev.Modifiers == 0 {
			switch ev.Key {
			case gdk.KEY_W:
				input.InputForward = true
			case gdk.KEY_S:
				input.InputBackward = true
			case gdk.KEY_A:
				input.InputLeft = true
			case gdk.KEY_D:
				input.InputRight = true
			case gdk.KEY_space:
				input.InputUp = true
			case gdk.KEY_Shift_L:
				input.InputDown = true
			case gdk.KEY_F1:
				input.InputF1 = true
			case gdk.KEY_F2:
				input.InputF2 = true
			}
		}
	} else if ev, ok := event.(ui.KeyUpEvent); ok {
		switch ev.Key {
		case gdk.KEY_W:
			input.InputForward = false
		case gdk.KEY_S:
			input.InputBackward = false
		case gdk.KEY_A:
			input.InputLeft = false
		case gdk.KEY_D:
			input.InputRight = false
		case gdk.KEY_space:
			input.InputUp = false
		case gdk.KEY_Shift_L:
			input.InputDown = false
		case gdk.KEY_F1:
			input.InputF1 = false
		case gdk.KEY_F2:
			input.InputF2 = false
		}
	} else if ev, ok := event.(ui.PointerMotionEvent); ok {
		// Store pointer position in screen space.

		x := ev.X
		y := ev.Y
		width := camera.WindowWidth
		height := camera.WindowHeight

		normalizedX := (float32(x) * 2 / float32(width)) - 1
		normalizedY := -(float32(y)*2/float32(height) - 1)

		input.PointerPosition = digimath.MakeVec2(normalizedX, normalizedY)

		// Store pointer ray in world space.

		v1 := digimath.MakeVec4(
			normalizedX, normalizedY, 0, 1,
		)

		v2 := digimath.MakeVec4(
			normalizedX, normalizedY, 1, 1,
		)

		// Revert the projection. The result are camera coordinates.
		v1 = camera.Noitcejorp.MulV(v1)
		v1 = v1.Scale(1 / v1.W())

		v2 = camera.Noitcejorp.MulV(v2)
		v2 = v2.Scale(1 / v2.W())

		t := node.
			GetData(&transform.TransformComponentKind).
			(transform.Transform)

		// Revert the camera. The result are world coordinates.
		aremac := InverseViewMatrix(t)

		v1 = aremac.MulV(v1)
		v2 = aremac.MulV(v2)

		ray := digimath.MakeRay(
			digimath.Vec3From4(v1), 
			digimath.Vec3From4(v2).Normalized(),
		)

		input.PointerRay = ray
	} else if ev, ok := event.(ui.PointerButtonDownEvent); ok {
		switch ev.Button {
		case ui.PointerButton1:
			input.LeftClicks += 1
		case ui.PointerButton2:
			input.RightClicks += 1
		}
	}

	ctx.OverwriteComponent(node.Entity, inputC.Update(input))
}

func InverseViewMatrix(t transform.Transform) digimath.Matrix44 {
	return t.ToMatrix()
}
