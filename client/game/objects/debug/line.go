package debug

import "github.com/CrimsonSarah/cto/client/digimath"

type DebugLineVertexData struct {
	Coords digimath.Vec3
}

// No need to make this pretty.
type DebugLine struct {
	Start DebugLineVertexData
	End   DebugLineVertexData
}

// This is what will get a transform and be drawn.
type DebugLines struct {
	Lines []DebugLine
	Color digimath.Vec3
}

func MakeDebugLine(
	start, end digimath.Vec3,
) DebugLine {
	return DebugLine{
		Start: DebugLineVertexData{
			Coords: start,
		},
		End: DebugLineVertexData{
			Coords: end,
		},
	}
}

func MakeDebugLines(color digimath.Vec3) DebugLines {
	return DebugLines{
		Lines: make([]DebugLine, 0),
		Color: color,
	}
}

func (d *DebugLines) Add(line DebugLine) {
	d.Lines = append(d.Lines, line)
}

// See `world/object.go`.
func (o DebugLines) Intersects(p digimath.Vec2) bool {
	// Implement if necessary.
	return false
}
