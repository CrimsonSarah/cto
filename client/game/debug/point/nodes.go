package point

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
)

var DebugPointRenderNodeKind = engine.MakeNodeKind(
	"DebugPointRender",
	&transform.TransformComponentKind,
	&DebugPointComponentKind,
)
