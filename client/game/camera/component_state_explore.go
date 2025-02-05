package camera

import (
	"github.com/CrimsonSarah/cto/client/engine"
)

var CameraStateExploreComponentKind = engine.MakeComponentKind("CameraStateExplore")

type CameraStateExplore struct{}

func (i CameraStateExplore) GetKind() *engine.ComponentKind {
	return &CameraStateExploreComponentKind
}

func MakeCameraStateExploreComponent() engine.Component {
	return engine.MakeComponent(CameraStateExplore{})
}
