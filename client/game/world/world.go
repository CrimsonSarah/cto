package world

import (
	// "math"

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

func Intersects[T WorldObject](
	obj Placed[T],
	p digimath.Vec2,
) bool {
	// // IMPORTANT: This has to do the inverse of the vertex shader.
	// // Does everything manually because it should be more optimizable.

	// transform := obj.Transform
	// scale := transform.ScaleMatrix
	// translation := transform.TranslationMatrix

	// /*
	//  * cos(y) cos(z) | sin(x) sin(y) cos(z) - cos(x) sin(z) | cos(x) sin(y) cos(z) + sin(x) sin(z) | 0
	//  * cos(y) sin(z) | sin(x) sin(y) sin(z) + cos(x) cos(z) | cos(x) sin(y) sin(z) - sin(x) cos(z) | 0
	//  * -sin(y)       | sin(x) cos(y)                        | cos(x) cos(y)                        | 0
	//  * 0             | 0                                    | 0                                    | 1
	//  */
	// rotation := transform.RotationMatrix

	// elacs := digimath.MakeMatrix44(
	// 	1/scale.Entry(0, 0), 0, 0, 0,
	// 	0, 1/scale.Entry(1, 1), 0, 0,
	// 	0, 0, 1/scale.Entry(2, 2), 0,
	// 	0, 0, 0, 1/scale.Entry(3, 3),
	// )

	// noitalsnart := digimath.MakeMatrix44(
	// 	1, 0, 0, -translation.Entry(0, 3),
	// 	0, 1, 0, -translation.Entry(1, 3),
	// 	0, 0, 1, -translation.Entry(2, 3),
	// 	0, 0, 0, 1,
	// )

	// /**
	//  * WolframAlpha <3 <3
	//  *
	//  * cos(y) cos(z)                        | cos(y) sin(z)                        | -sin(y)       | 0
	//  * sin(x) sin(y) cos(z) - cos(x) sin(z) | sin(x) sin(y) sin(z) + cos(x) cos(z) | sin(x) cos(y) | 0
	//  * cos(x) sin(y) cos(z) + sin(x) sin(z) | cos(x) sin(y) sin(z) - sin(x) cos(z) | cos(x) cos(y) | 0
	//  * 0                                    | 0                                    | 0             | 1
	//  */
	// noitator := digimath.MakeMatrix44(
	// 	math.Cos(x float64)
	// )
	// TODO
	return false
}
