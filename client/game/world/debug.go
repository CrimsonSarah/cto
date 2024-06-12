package world

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/objects/debug"
)

type WorldDebug struct {
	// Slice of slices. Each inner slice gets a potentially
	// different Transform. Will be more efficient for
	// drawing complex frames, though I'm not sure that
	// matters...
	LineGroups []*Placed[debug.DebugLines]

	Points []*Placed[debug.DebugPoint]
}

// Start, end in local coordinates.
// Transform still applies.
func (w *World) AddLine(
	start, end digimath.Vec3,
	color digimath.Vec3,
) *Placed[debug.DebugLines] {
	lines := debug.MakeDebugLines(color)
	line := debug.MakeDebugLine(start, end)
	lines.Add(line)

	placed := MakePlacedDefault(&lines, w)

	w.Debug.LineGroups = append(w.Debug.LineGroups, &placed)
	return &placed
}

// `coords` shall be a slice where each item is a tuple with
// start and end, in that order.
// Start, end in local coordinates.
// Transform still applies.
func (w *World) AddLines(
	coords [][2]digimath.Vec3,
	color digimath.Vec3,
) *Placed[debug.DebugLines] {
	lines := debug.MakeDebugLines(color)

	for _, coord := range coords {
		start := coord[0]
		end := coord[1]

		line := debug.MakeDebugLine(start, end)
		lines.Add(line)
	}

	placed := MakePlacedDefault(&lines, w)

	w.Debug.LineGroups = append(w.Debug.LineGroups, &placed)
	return &placed
}

// Start, end in clip space.
// Transform still applies.
func (w *World) AddClipLine(
	start, end digimath.Vec2,
	color digimath.Vec3,
) *Placed[debug.DebugLines] {
	start4 := digimath.MakeVec4(
		start.X(), start.Y(), 0, 1,
	)

	start4 = w.Noitcejorp.MulV(start4)

	end4 := digimath.MakeVec4(
		end.X(), end.Y(), 0, 1,
	)

	end4 = w.Noitcejorp.MulV(end4)

	return w.AddLine(
		digimath.Vec3From4(start4),
		digimath.Vec3From4(end4),
		color,
	)
}

// Start, end in local coordinates.
// Transform still applies.
func (w *World) AddPoint(
	center digimath.Vec3,
	color digimath.Vec3,
) *Placed[debug.DebugPoint] {
	point := debug.MakeDebugPoint(center, color)
	placed := MakePlacedDefault(&point, w)

	w.Debug.Points = append(w.Debug.Points, &placed)
	return &placed
}

// Start, end in clip space.
// Transform still applies.
func (w *World) AddClipPoint(
	center digimath.Vec2,
	color digimath.Vec3,
) *Placed[debug.DebugPoint] {
	center4 := digimath.MakeVec4(
		center.X(), center.Y(), 0, 1,
	)

	center4 = w.Noitcejorp.MulV(center4)

	return w.AddPoint(
		digimath.Vec3From4(center4),
		color,
	)
}
