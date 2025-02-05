package camera

import (
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
)

var CameraInputNodeKind = engine.MakeNodeKind(
	"CameraInput",
	&CameraInputComponentKind,
	&CameraCameraComponentKind,
	&transform.TransformComponentKind,
)

var CameraExploreNodeKind = engine.MakeNodeKind(
	"CameraExplore",
	&CameraStateExploreComponentKind,
	&CameraInputComponentKind,
	&transform.TransformComponentKind,
)

var CameraDefaultNodeKind = engine.MakeNodeKind(
	"CameraDefault",
	&CameraStateDefaultComponentKind,
	&CameraInputComponentKind,
	&transform.TransformComponentKind,
)
