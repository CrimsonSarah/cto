package debug

import "github.com/CrimsonSarah/cto/client/digimath"

type DebugPointVertexData struct {
	Coords [4]digimath.Vec3
}

type DebugPoint struct {
	Center digimath.Vec3
	Color  digimath.Vec3
}

var DebugPointVertexIndices = [6]uint32{
	0, 1, 2, 2, 3, 0,
}

func MakeDebugPoint(
	center, color digimath.Vec3,
) DebugPoint {
	return DebugPoint{
		Center: center,
		Color:  color,
	}
}

// See `world/object.go`.
func (o DebugPoint) Intersects(p digimath.Vec2) bool {
	// Implement if necessary.
	return false
}
