package camera

import (
	"github.com/CrimsonSarah/cto/client/engine"
)

var CameraStateDefaultComponentKind = engine.MakeComponentKind("CameraStateDefault")

type CameraStateDefault struct{}

func (i CameraStateDefault) GetKind() *engine.ComponentKind {
	return &CameraStateDefaultComponentKind
}

func MakeCameraStateDefaultComponent() engine.Component {
	return engine.MakeComponent(CameraStateDefault{})
}
