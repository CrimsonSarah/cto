package world

import (
	"fmt"

	"github.com/CrimsonSarah/cto/client/digimath"
)

// This file assumes there is only 1 world and describes mostly a
// coordinate system.

// TODO: Is this optimal?
// These are combined on the GPU
type Transform struct {
	ScaleMatrix       digimath.Matrix44
	RotationMatrix    digimath.Matrix44
	TranslationMatrix digimath.Matrix44
}

func (t *Transform) Scale(amount float32) {
	t.ScaleMatrix = digimath.Matrix44Scale(amount).Mul(
		t.ScaleMatrix,
	)
}

func (t *Transform) Translate(amount digimath.Vec3) {
	t.TranslationMatrix = digimath.Matrix44Translate(amount).Mul(
		t.TranslationMatrix,
	)
}

func (t *Transform) TranslateX(amount float32) {
	t.TranslationMatrix = digimath.Matrix44Translate(
		digimath.MakeVec3(amount, 0, 0),
	).Mul(
		t.TranslationMatrix,
	)
}

func (t *Transform) TranslateY(amount float32) {
	t.TranslationMatrix = digimath.Matrix44Translate(
		digimath.MakeVec3(0, amount, 0),
	).Mul(
		t.TranslationMatrix,
	)

}

func (t *Transform) TranslateZ(amount float32) {
	t.TranslationMatrix = digimath.Matrix44Translate(
		digimath.MakeVec3(0, 0, amount),
	).Mul(
		t.TranslationMatrix,
	)
}

func (t *Transform) RotateX(amount float32) {
	t.RotationMatrix = digimath.Matrix44RotateX(amount).Mul(
		t.RotationMatrix,
	)
}

func (t *Transform) RotateY(amount float32) {
	t.RotationMatrix = digimath.Matrix44RotateY(amount).Mul(
		t.RotationMatrix,
	)
}

func (t *Transform) RotateZ(amount float32) {
	t.RotationMatrix = digimath.Matrix44RotateZ(amount).Mul(
		t.RotationMatrix,
	)
}

func (t *Transform) GetPosition() digimath.Vec3 {
	m := t.TranslationMatrix
	return digimath.MakeVec3(
		m.Entry(0, 3),
		m.Entry(1, 3),
		m.Entry(2, 3),
	)
}

// For debugging
func (t *Transform) Format() string {
	return fmt.Sprintf(
		""+
			"Scale:\n%s\n\n"+
			"Rotation:\n%s\n\n"+
			"Translation:\n%s",
		t.ScaleMatrix.Format(),
		t.RotationMatrix.Format(),
		t.TranslationMatrix.Format(),
	)
}

// A thing in this world.
type Placed[T any] struct {
	Obj       *T
	Id        uint32
	Transform Transform
}

var currentId uint32 = 0

func MakePlacedDefault[T any](obj *T) Placed[T] {
	id := currentId
	currentId += 1

	return Placed[T]{
		Obj: obj,
		Id:  id,
		Transform: Transform{
			ScaleMatrix:       digimath.Matrix44Id(),
			RotationMatrix:    digimath.Matrix44Id(),
			TranslationMatrix: digimath.Matrix44Id(),
		},
	}
}
