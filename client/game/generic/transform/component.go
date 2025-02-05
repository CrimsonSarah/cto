package transform

import (
	"fmt"
	m "math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
)

var TransformComponentKind = engine.MakeComponentKind("Transform")

// Easy to manipulate. Becomes a matrix for the GPU.
type Transform struct {
	ScaleFactor float32
	// TODO: Use rotors / quaternions.
	Rotation digimath.Vec3 // Radians
	Position digimath.Vec3
}

func MakeTransform() Transform {
	return Transform{
		ScaleFactor: 1,
		Rotation:    digimath.Vec3Zero,
		Position:    digimath.Vec3Zero,
	}
}

func (t *Transform) Scale(amount float32) {
	t.ScaleFactor *= amount
}

func (t *Transform) Translate(amount digimath.Vec3) {
	t.Position = t.Position.Add(amount)
}

func (t *Transform) TranslateX(amount float32) {
	t.Translate(digimath.MakeVec3(amount, 0, 0))
}

func (t *Transform) TranslateY(amount float32) {
	t.Translate(digimath.MakeVec3(0, amount, 0))
}

func (t *Transform) TranslateZ(amount float32) {
	t.Translate(digimath.MakeVec3(0, 0, amount))
}

func (t *Transform) Rotate(radians digimath.Vec3) {
	t.Rotation = t.Rotation.Add(radians)
}

func (t *Transform) RotateX(radians float32) {
	t.Rotate(digimath.MakeVec3(radians, 0, 0))
}

func (t *Transform) RotateY(radians float32) {
	t.Rotate(digimath.MakeVec3(0, radians, 0))
}

func (t *Transform) RotateZ(radians float32) {
	t.Rotate(digimath.MakeVec3(0, 0, radians))
}

func (t *Transform) GetPosition() digimath.Vec3 {
	return t.Position
}

func (t *Transform) GetDistanceFromCamera() float32 {
	// No camera for now.
	return -t.Position.Z()
}

func (t *Transform) ToMatrix() digimath.Matrix44 {
	// Scale
	//
	//   s | 0 | 0 | 0
	//   0 | s | 0 | 0
	//   0 | 0 | s | 0
	//   0 | 0 | 0 | s
	//
	// Translation
	//
	//   1 | 0 | 0 | x
	//   0 | 1 | 0 | y
	//   0 | 0 | 1 | z
	//   0 | 0 | 0 | 1
	//
	// Rotation (rx, ry and rz are independent variables)
	//
	//   cos(ry) cos(rz) | sin(rx) sin(ry) cos(rz) - cos(rx) sin(rz) | cos(rx) sin(ry) cos(rz) + sin(rx) sin(rz) | 0
	//   cos(ry) sin(rz) | sin(rx) sin(ry) sin(rz) + cos(rx) cos(rz) | cos(rx) sin(ry) sin(rz) - sin(rx) cos(rz) | 0
	//   -sin(ry)        | sin(rx) cos(ry)                           | cos(rx) cos(ry)                           | 0
	//   0               | 0                                         | 0                                         | 1
	//
	// Aggregate
	//
	// s * (
	//   cos(rz) cos(ry) | cos(rz) sin(ry) sin(rx) - sin(rz) cos(rx) | cos(rz) sin(ry) cos(rx) + sin(rz) sin(rx) | x
	//   sin(rz) cos(ry) | sin(rz) sin(ry) sin(rx) + cos(rz) cos(rx) | sin(rz) sin(ry) cos(rx) - cos(rz) sin(rx) | y
	//   -sin(ry)        | cos(ry) sin(rx)                           | cos(ry) cos(rx)                           | z
	//   0               | 0                                         | 0                                         | 1
	// )
	//
	// Last column is rescaled to keep translation absolute.

	s := t.ScaleFactor
	p := t.Position
	r := t.Rotation

	cos := func(x float32) float32 { return float32(m.Cos(float64(x))) }
	sin := func(x float32) float32 { return float32(m.Sin(float64(x))) }

	cosx := cos(r.X())
	cosy := cos(r.Y())
	cosz := cos(r.Z())
	sinx := sin(r.X())
	siny := sin(r.Y())
	sinz := sin(r.Z())

	e11 := s * (cosz * cosy)
	e12 := s * (cosz*siny*sinx - sinz*cosx)
	e13 := s * (cosz*siny*cosx + sinz*sinx)
	e14 := p.X()

	e21 := s * (sinz * cosy)
	e22 := s * (sinz*siny*sinx + cosz*cosx)
	e23 := s * (sinz*siny*cosx - cosz*sinx)
	e24 := p.Y()

	e31 := s * (-siny)
	e32 := s * (cosy * sinx)
	e33 := s * (cosy * cosx)
	e34 := p.Z()

	var e41, e42, e43 float32 = 0, 0, 0
	e44 := float32(1)

	return digimath.MakeMatrix44(
		e11, e12, e13, e14,
		e21, e22, e23, e24,
		e31, e32, e33, e34,
		e41, e42, e43, e44,
	)
}

// TODO: Benchmark.
func (t *Transform) ToMatrix2() digimath.Matrix44 {
	s := t.ScaleFactor
	p := t.Position
	r := t.Rotation

	mat := digimath.Matrix44Id()
	mat = digimath.Matrix44RotateX(r.X()).Mul(mat)
	mat = digimath.Matrix44RotateY(r.Y()).Mul(mat)
	mat = digimath.Matrix44RotateZ(r.Z()).Mul(mat)
	mat = digimath.Matrix44Translate(p).Mul(mat)
	mat = digimath.Matrix44Scale(s).Mul(mat)

	return mat
}

// TODO: Do the transposition thing.
func (t *Transform) ToMatrixInverse() digimath.Matrix44 {
	matrix := t.ToMatrix()
	return matrix.Inverse0001()
}

// For debugging.
func (t *Transform) Format() string {
	return fmt.Sprintf(
		""+
			"Scale:\n%f\n\n"+
			"Rotation:\n%v\n\n"+
			"Translation:\n%v\n\n",
		t.ScaleFactor,
		t.Rotation,
		t.Position,
	)
}

func (t Transform) GetForward() digimath.Vec3 {
	return digimath.
		Matrix33Rotate(t.Rotation).
		MulV(digimath.MakeVec3(0, 0, -1))
}

// Removes the Y component.
func (t Transform) GetForwardFlat() digimath.Vec3 {
	forward := t.GetForward()
	forward.SetY(0)
	return forward.Normalized()
}

func (t Transform) ApplyToRay(ray digimath.Ray, pivot digimath.Vec3) digimath.Ray {
	toOrigin := ray.Origin.Sub(pivot)

	newToOrigin := digimath.Vec3From4(
		digimath.Matrix44Translate(t.Position).MulV(toOrigin.AsPoint()),
	)

	newOrigin := newToOrigin.Add(pivot)

	newDirection := digimath.Vec3From4(
		digimath.Matrix44Rotate3(t.Rotation).MulV(ray.Direction.AsDirection()),
	)

	return digimath.MakeRay(newOrigin, newDirection)
}

func (t Transform) ApplyToPlane(plane digimath.Plane, pivot digimath.Vec3) digimath.Plane {
	panic("Applying transform to plano not implemented.")
}

func (t Transform) ApplyToRect(rect digimath.Rect, pivot digimath.Vec3) digimath.Rect {
	toBottomLeft := rect.GetBottomLeft().Sub(pivot)

	m := t.ToMatrix()
	newToBottomLeft := digimath.Vec3From4(m.MulV(toBottomLeft.AsPoint()))
	newSide1 := digimath.Vec3From4(m.MulV(rect.Side1.AsDirection()))
	newSide2 := digimath.Vec3From4(m.MulV(rect.Side2.AsDirection()))

	newBottomLeft := newToBottomLeft.Add(pivot)
	return digimath.MakeRect(newBottomLeft, newSide1, newSide2)
}

func (t Transform) GetKind() *engine.ComponentKind {
	return &TransformComponentKind
}

func MakeTransformComponent() engine.Component {
	return engine.MakeComponent(MakeTransform())
}
