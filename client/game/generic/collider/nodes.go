package collider

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
)

var ColliderNodeKind = engine.MakeNodeKind(
	"Collider",
	&ColliderComponentKind,
	&transform.TransformComponentKind,
)
