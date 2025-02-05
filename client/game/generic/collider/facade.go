package collider

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/global"
)

func AddCollisionSystems(
	e *engine.Engine[global.GameGlobals],
) {
	e.AddSystem(ColliderSystem, 0)
}
