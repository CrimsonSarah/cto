package world

import (
	// "math"

	"fmt"

	"github.com/CrimsonSarah/cto/client/digimath"
)

type World struct {
	Projection digimath.Matrix44
	Noitcejorp digimath.Matrix44
}

func MakeWorld(width, height int) World {
	projection := GetProjection(
		float32(width),
		float32(height),
	)

	noitcejorp := GetNoitcejorp(projection)

	return World{
		Projection: projection,
		Noitcejorp: noitcejorp,
	}
}

func (w *World) Configure(width, height int) {
	w.Projection = GetProjection(
		float32(width),
		float32(height),
	)
	w.Noitcejorp = GetNoitcejorp(w.Projection)
}

// `p` should be in the range [-1, +1], where X -> -1 is left and
// Y -> -1 is down.
func Intersects[T WorldObject](
	obj Placed[T],
	p digimath.Vec2,
) bool {

	index := digimath.MakeVec4(
		p.X(), p.Y(), 0, 1,
	)

	fmt.Println("")
	// Revert depth.
	// TODO: Rotation should probably have an effect.
	depth := obj.Transform.GetDistanceFromCamera()
	index = index.Scale(depth)
	fmt.Printf("Index 0 %v\n", index)

	// We need to find Z such that, after multiplying by the projection
	// inverse, W = 1.
	// This is just how the projection computes Z.
	objPos := obj.Transform.GetPosition()

	projectedZ :=
		objPos.Z()*(-2/(PROJECTION_FAR-PROJECTION_NEAR)) -
			((PROJECTION_FAR + PROJECTION_NEAR) / (PROJECTION_FAR - PROJECTION_NEAR))
	index.SetZ(projectedZ)

	// With the correct Z and W in place, invert the projection.
	// The result are world coordinates.
	index = obj.World.Noitcejorp.MulV(index)

	fmt.Printf("Index 1 %v\n", index)

	// index = digimath.MakeVec4(
	// 	index.X()-objPos.X(),
	// 	index.Y()-objPos.Y(),
	// 	index.Z()-objPos.Z(),
	// 	index.W(),
	// )

	transInv := obj.Transform.ToMatrixInverse()
	fmt.Printf("Transform inverse\n%s\n", transInv.Format())

	// Revert the transform. The result are local coordinates.
	index = obj.Transform.ToMatrixInverse().MulV(index)

	fmt.Printf("Index 2 %v\n", index)
	fmt.Printf("ObjPos %v\n", objPos)
	fmt.Println("")

	return (*obj.Obj).Intersects(
		digimath.MakeVec2(index.X(), index.Y()),
	)
}
