package sprite

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
)

var SpriteRenderNodeKind = engine.MakeNodeKind(
	"SpriteRender",
	&transform.TransformComponentKind,
	&SpriteComponentKind,
)
