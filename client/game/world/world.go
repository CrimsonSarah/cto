package world

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/objects/debug"
)

type World struct {
	WindowWidth  int
	WindowHeight int

	Camera Transform

	Projection digimath.Matrix44
	Noitcejorp digimath.Matrix44

	Debug WorldDebug
}

func MakeWorld(width, height int) World {
	projection := GetProjection(
		float32(width),
		float32(height),
	)

	noitcejorp := GetNoitcejorp(projection)

	// Debug.
	lineGroups := make([]*Placed[debug.DebugLines], 0)
	points := make([]*Placed[debug.DebugPoint], 0)

	return World{
		Camera:     MakeTransform(),
		Projection: projection,
		Noitcejorp: noitcejorp,
		Debug: WorldDebug{
			LineGroups: lineGroups,
			Points:     points,
		},
	}
}

func (w *World) Configure(width, height int) {
	w.WindowWidth = width
	w.WindowHeight = height

	w.Projection = GetProjection(
		float32(width),
		float32(height),
	)
	w.Noitcejorp = GetNoitcejorp(w.Projection)
}

func (w *World) CameraMatrix() digimath.Matrix44 {
	return w.Camera.ToMatrixInverse()
}

func (w *World) AremacMatrix() digimath.Matrix44 {
	return w.Camera.ToMatrix()
}

func (w *World) CameraForward() digimath.Vec3 {
	return digimath.Vec3From4(w.AremacMatrix().MulV(
		digimath.MakeVec3(0, 0, -1).AsDirection(),
	))
}

// Removes the Y component.
func (w *World) CameraForwardFlat() digimath.Vec3 {
	forward := w.CameraForward()
	forward.SetY(0)
	return forward.Normalized()
}

// `p` should be in the range [-1, +1], where X -> -1 is left and
// Y -> -1 is down.
// Currently only works for quads. TODO: Organize.
func Intersects[T WorldObject](
	obj Placed[T],
	p digimath.Vec2,
) bool {
	ray0 := digimath.MakeVec4(
		p.X(), p.Y(), 0, 1,
	)

	ray1 := digimath.MakeVec4(
		p.X(), p.Y(), 1, 1,
	)

	// Revert the projection. The result are camera coordinates.
	ray0 = obj.World.Noitcejorp.MulV(ray0)
	ray0 = ray0.Scale(1 / ray0.W())

	ray1 = obj.World.Noitcejorp.MulV(ray1)
	ray1 = ray1.Scale(1 / ray1.W())

	// Revert the camera. The result are world coordinates.
	aremac := obj.World.AremacMatrix()
	ray0 = aremac.MulV(ray0)
	ray1 = aremac.MulV(ray1)

	// Unnormalized. Don't think it's important here.
	line := digimath.MakeLine(
		digimath.Vec3From4(ray0), digimath.Vec3From4(ray1),
	)

	rotation := obj.Transform.Rotation
	normal := digimath.Matrix44RotateX(rotation.X()).
		Mul(digimath.Matrix44RotateY(rotation.Y())).
		MulV(digimath.MakeVec4(0, 0, 1, 0))

	d := -normal.Dot(obj.Transform.Position.AsPoint())

	plane := digimath.MakePlane(digimath.Vec3From4(normal), d)
	intersects, p2 := digimath.IntersectLinePlane(line, plane)

	if !intersects {
		return false
	}

	transInv := obj.Transform.ToMatrixInverse()

	// Revert the transform. The result are local coordinates.
	p2 = digimath.Vec3From4(transInv.MulV(p2.AsPoint()))

	return (*obj.Obj).Intersects(
		digimath.MakeVec2(p2.X(), p2.Y()),
	)
}
