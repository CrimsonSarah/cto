package camera

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/global"
)

var CameraStates = []*engine.ComponentKind{
	&CameraStateDefaultComponentKind,
	&CameraStateExploreComponentKind,
}

func CameraToDefaultState(
	e engine.IEngine[global.GameGlobals],
	entity *engine.Entity,
) {
	e.SwitchComponents(entity, CameraStates, &CameraStateDefaultComponentKind)
}

func CameraToExploreState(
	e engine.IEngine[global.GameGlobals],
	entity *engine.Entity,
) {
	e.SwitchComponents(entity, CameraStates, &CameraStateExploreComponentKind)
}
