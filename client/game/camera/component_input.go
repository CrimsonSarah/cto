package camera

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
)

var CameraInputComponentKind = engine.MakeComponentKind("CameraInput")

// TODO: Maybe use a bitfield?..
type CameraInput struct {
	InputForward  bool
	InputBackward bool
	InputUp       bool
	InputDown     bool
	InputLeft     bool
	InputRight    bool

	InputF1 bool
	InputF2 bool

	// In [-1, +1] range for each axis.
	PointerPosition digimath.Vec2

	// In world coordinates.
	PointerRay digimath.Ray

	// How many clicks there are to process.
	LeftClicks uint
	RightClicks uint
}

func MakeCameraInput() CameraInput {
	return CameraInput{}
}

// Get and consume.
func (i *CameraInput) Get1LeftClick() bool {
	if i.LeftClicks >= 1 {
		i.LeftClicks -= 1	
		return true
	}

	return false
}

// Get and consume.
func (i *CameraInput) Get1RightClick() bool {
	if i.LeftClicks >= 1 {
		i.LeftClicks -= 1	
		return true
	}

	return false
}

// Get2Clicks() would require more work.

func (i CameraInput) GetKind() *engine.ComponentKind {
	return &CameraInputComponentKind
}

func MakeCameraInputComponent() engine.Component {
	return engine.MakeComponent(MakeCameraInput())
}
