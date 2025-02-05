package engine

// TODO: Ponder on the usage of pointers in this.

import (
	"github.com/CrimsonSarah/cto/client/digidata"
	"github.com/CrimsonSarah/cto/client/game/world"
)

// Interface implemented by both Engine and EngineContext.
// Ideally this would contain every operation in Engine.
type IEngine[Global any] interface {
	GetGlobals() *Global
	AddEntity() Entity
	RemoveEntity(entity Entity)
	GetComponents(kind ComponentKindId) *digidata.SparseArray[Component]
	AddComponent(entity *Entity, component Component)
	EnableComponent(entity *Entity, kind *ComponentKind)
	RemoveComponent(entity *Entity, kind *ComponentKind)
	DisableComponent(entity *Entity, kind *ComponentKind)
	OverwriteComponent(entity *Entity, component Component)
	SwitchComponents(entity *Entity, options []*ComponentKind, current *ComponentKind)
	GetNodes(kind NodeKindId) []Node
	AddEvent(event Event)
}

type componentIdTuple struct {
	Kind   *ComponentKind
	Entity *Entity
}

type componentTuple struct {
	Component Component
	Entity    *Entity
}

type componentSwitchTuple struct {
	Entity  *Entity
	Options []*ComponentKind
	Current *ComponentKind
}

// Stores operations to be applied at the end of the system processing
// step. Allows for sane parallelization.
// If there is something in Engine that's not in EngineContext, it should
// be more or less straightforwared to add.
//
// NOTE: If you add anything to this, remember to clean it up in Clear.
// TODO: System manipulation.
type EngineContext[Global any] struct {
	engine *Engine[Global]

	addEntities    []Entity
	removeEntities []Entity

	addComponents       []componentTuple
	enableComponents    []componentIdTuple
	removeComponents    []componentIdTuple
	disableComponents   []componentIdTuple
	overwriteComponents []componentTuple
	switchComponents    []componentSwitchTuple

	addEvents []Event
}

func MakeEngineContext[Global any](e *Engine[Global]) EngineContext[Global] {
	return EngineContext[Global]{
		engine:           e,
		addEntities:      nil,
		removeEntities:   nil,
		addComponents:    nil,
		removeComponents: nil,
	}
}

func (e *EngineContext[Global]) GetWorld() *world.World {
	return e.engine.World
}

func (e *EngineContext[Global]) GetGlobals() *Global {
	return e.engine.GetGlobals()
}

func (e *EngineContext[Global]) AddEntity() Entity {
	entity := MakeEntity()
	e.addEntities = append(e.addEntities, entity)
	return entity
}

func (e *EngineContext[Global]) RemoveEntity(entity Entity) {
	e.removeEntities = append(e.removeEntities, entity)
}

func (e *EngineContext[Global]) GetComponents(kind ComponentKindId) *digidata.SparseArray[Component] {
	return e.engine.GetComponents(kind)
}

func (e *EngineContext[Global]) AddComponent(entity *Entity, component Component) {
	e.addComponents = append(e.addComponents, componentTuple{
		Component: component,
		Entity:    entity,
	})
}

func (e *EngineContext[Global]) EnableComponent(entity *Entity, kind *ComponentKind) {
	e.enableComponents = append(e.enableComponents, componentIdTuple{
		Kind:   kind,
		Entity: entity,
	})
}

func (e *EngineContext[Global]) RemoveComponent(entity *Entity, kind *ComponentKind) {
	e.removeComponents = append(e.removeComponents, componentIdTuple{
		Kind:   kind,
		Entity: entity,
	})
}

func (e *EngineContext[Global]) DisableComponent(entity *Entity, kind *ComponentKind) {
	e.disableComponents = append(e.disableComponents, componentIdTuple{
		Kind:   kind,
		Entity: entity,
	})
}

func (e *EngineContext[Global]) OverwriteComponent(entity *Entity, component Component) {
	e.overwriteComponents = append(e.overwriteComponents, componentTuple{
		Component: component,
		Entity:    entity,
	})
}

func (e *EngineContext[Global]) SwitchComponents(entity *Entity, options []*ComponentKind, current *ComponentKind) {
	e.switchComponents = append(e.switchComponents, componentSwitchTuple{
		Entity:  entity,
		Options: options,
		Current: current,
	})
}

func (e *EngineContext[Global]) GetNodes(kind NodeKindId) []Node {
	return e.engine.GetNodes(kind)
}

func (e *EngineContext[Global]) AddEvent(event Event) {
	e.addEvents = append(e.addEvents, event)
}

func (e *EngineContext[Global]) Clean() {
	e.addEntities = e.addEntities[:0]
	e.removeEntities = e.removeEntities[:0]
	e.addComponents = e.addComponents[:0]
	e.enableComponents = e.enableComponents[:0]
	e.removeComponents = e.removeComponents[:0]
	e.disableComponents = e.disableComponents[:0]
	e.overwriteComponents = e.overwriteComponents[:0]
	e.switchComponents = e.switchComponents[:0]
	e.addEvents = e.addEvents[:0]
}
