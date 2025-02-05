package sprite

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/global"
)

func AddSpriteSystems(
	e *engine.Engine[global.GameGlobals],
) {
	e.AddSystem(SpriteRenderSystem, 1)
}
