package world

import "github.com/CrimsonSarah/cto/client/digimath"

// This file assumes there is only 1 world and describes mostly a
// coordinate system.

type Transform struct {
	digimath.Matrix44
}

func (t *Transform) Scale(amount float32) {
	t.Matrix44 = digimath.Matrix44Scale(amount).Mul(
		t.Matrix44,
	)
}

func (t *Transform) Translate(amount digimath.Vec3) {
	t.Matrix44 = digimath.Matrix44Translate(amount).Mul(
		t.Matrix44,
	)
}

func (t *Transform) TranslateX(amount float32) {
	t.Matrix44 = digimath.Matrix44Translate(
		digimath.MakeVec3(amount, 0, 0),
	).Mul(
		t.Matrix44,
	)
}

func (t *Transform) TranslateY(amount float32) {
	t.Matrix44 = digimath.Matrix44Translate(
		digimath.MakeVec3(0, amount, 0),
	).Mul(
		t.Matrix44,
	)

}

func (t *Transform) TranslateZ(amount float32) {
	t.Matrix44 = digimath.Matrix44Translate(
		digimath.MakeVec3(0, 0, amount),
	).Mul(
		t.Matrix44,
	)
}

func (t *Transform) RotateX(amount float32) {
	t.Matrix44 = digimath.Matrix44RotateX(amount).Mul(
		t.Matrix44,
	)
}

func (t *Transform) RotateY(amount float32) {
	t.Matrix44 = digimath.Matrix44RotateY(amount).Mul(
		t.Matrix44,
	)
}

func (t *Transform) RotateZ(amount float32) {
	t.Matrix44 = digimath.Matrix44RotateZ(amount).Mul(
		t.Matrix44,
	)
}

func (t *Transform) GetPosition() digimath.Vec3 {
	return digimath.MakeVec3(t.Entry(0, 3), t.Entry(1, 3), t.Entry(2, 3))
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
		Obj:       obj,
		Id:        id,
		Transform: Transform{digimath.Matrix44Id()},
	}
}

func MakePlaced[T any](obj *T, transform digimath.Matrix44) Placed[T] {
	id := currentId
	currentId += 1

	return Placed[T]{
		Obj:       obj,
		Id:        id,
		Transform: Transform{transform},
	}
}
