package engine

import (
	"github.com/CrimsonSarah/cto/client/game/world"
)

type SystemImpl[Global any] interface {
	EventKinds() []EventKindId
	NodeKind() *NodeKind

	Init(ctx *EngineContext[Global], event any, node Node)
	Handle(ctx *EngineContext[Global], event any, node Node)
}

type InitContext struct {
	World *world.World

	Width  int
	Height int
}

// Satisfies `digidata.Comparable`.
type System[Global any] struct {
	Value uint64
	Name  string

	Impl SystemImpl[Global]
}

var systemValueCounter = 0

func MakeSystem[Global any](name string, impl SystemImpl[Global]) System[Global] {
	counter := systemValueCounter
	systemValueCounter += 1

	return System[Global]{
		Value: Primes[counter],
		Name:  name,
		Impl:  impl,
	}
}

type SystemId int64

func (s System[Global]) Id() SystemId {
	return SystemId(s.Value)
}

func (s System[Global]) EventKinds() []EventKindId {
	return s.Impl.EventKinds()
}

func (s System[Global]) NodeKind() *NodeKind {
	return s.Impl.NodeKind()
}

func (s System[Global]) Init(ctx *EngineContext[Global], event any, node Node) {
	s.Impl.Init(ctx, event, node)
}

func (s System[Global]) Handle(ctx *EngineContext[Global], event any, node Node) {
	s.Impl.Handle(ctx, event, node)
}

func (c1 System[Global]) Equals(c2 System[Global]) bool {
	return c1.Value == c2.Value
}
