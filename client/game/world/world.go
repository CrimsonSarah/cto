package world

import "github.com/CrimsonSarah/cto/client/digimath"

// This file assumes there is only 1 world and describes mostly a
// coordinate system.

// A thing in this world.
type Placed[T any] struct {
	Obj       *T
	Id        uint32
	Transform digimath.Matrix44
}

var currentId uint32 = 0

func MakePlacedDefault[T any](obj *T) Placed[T] {
	id := currentId
	currentId += 1

	return Placed[T]{
		Obj:       obj,
		Id:        id,
		Transform: digimath.Matrix44Id(),
	}
}

func MakePlaced[T any](obj *T, transform digimath.Matrix44) Placed[T] {
	id := currentId
	currentId += 1

	return Placed[T]{
		Obj:       obj,
		Id:        id,
		Transform: transform,
	}
}
