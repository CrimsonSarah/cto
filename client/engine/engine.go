package engine

import (
	"github.com/CrimsonSarah/cto/client/digidata"
	"github.com/CrimsonSarah/cto/client/game/world"
)

// If a field in Global is meant to be modified in a system, consider
// making it "thread"-safe in case we do decide to run systems
// concurrently.
type Engine[Global any] struct {
	World   *world.World
	Globals Global

	// Stores operations to later be applied simultaneously-ish.
	// Is reused between runs to save on allocations.
	Context EngineContext[Global]

	Entities   digidata.SparseArray[Entity]
	Components map[ComponentKindId]*digidata.SparseArray[Component]
	Nodes      map[NodeKindId]*digidata.SparseArray[StoredNode]

	Systems map[EventKindId]*digidata.PrioritizedUnorderedArray[System[Global]]
	Events  digidata.Queue[Event]
}

func NewEngine[Global any](w *world.World) *Engine[Global] {
	e := Engine[Global]{
		World:      w,
		Entities:   digidata.MakeSparseArray[Entity](),
		Components: make(map[ComponentKindId]*digidata.SparseArray[Component]),
		Nodes:      make(map[NodeKindId]*digidata.SparseArray[StoredNode]),
		Systems:    make(map[EventKindId]*digidata.PrioritizedUnorderedArray[System[Global]]),
		Events:     digidata.MakeQueue[Event](),
	}

	e.Context = MakeEngineContext(&e)
	return &e
}

// BEWARE! The pointer is valid only until another operation is
// performed on entities!
func (e *Engine[Global]) AddNewEntity() *Entity {
	entity := MakeEntity()
	return e.AddEntity(entity)
}

// BEWARE! The pointer is valid only until another operation is
// performed on entities!
func (e *Engine[Global]) AddEntity(entity Entity) *Entity {
	e.Entities.Upsert(int(entity.Id), entity)

	ptr, ok := e.Entities.GetPtr(int(entity.Id))

	if !ok {
		panic("Could not find entity immediately after insertion")
	}

	e.AddEvent(Event{
		KindId: EntityCreatedEventKindId,
		Data: EntityCreatedEvent{
			EntityId: entity.Id,
		},
	})

	return ptr
}

// TODO: Remove associated components and stored nodes.
func (e *Engine[Global]) RemoveEntity(entity Entity) {
	e.Entities.Remove(int(entity.Id))
}

// TODO: Create one for each kind at startup and don't check every
// time.
func (e *Engine[Global]) GetComponents(kind ComponentKindId) *digidata.SparseArray[Component] {
	components, ok := e.Components[kind]

	if !ok {
		c := digidata.MakeSparseArray[Component]()
		components = &c

		e.Components[kind] = components
	}

	return components
}

// TODO: ^
func (e *Engine[Global]) GetStoredNodes(kind NodeKindId) *digidata.SparseArray[StoredNode] {
	nodes, ok := e.Nodes[kind]

	if !ok {
		n := digidata.MakeSparseArray[StoredNode]()
		nodes = &n

		e.Nodes[kind] = nodes
	}

	return nodes
}

func (e *Engine[Global]) GetSystems(kind EventKindId) *digidata.PrioritizedUnorderedArray[System[Global]] {
	systems, ok := e.Systems[kind]

	if !ok {
		s := digidata.MakePrioritizedUnorderedArray[System[Global]]()
		systems = &s

		e.Systems[kind] = systems
	}

	return systems
}

func (e *Engine[Global]) GetGlobals() *Global {
	return &e.Globals
}

// Will set the component's ID.
func (e *Engine[Global]) AddComponent(entity *Entity, component Component) {
	components := e.GetComponents(component.GetKind().Id)

	component.EntityId = entity.Id
	components.Upsert(int(entity.Id), component)

	e.EnableComponent(entity, component.GetKind())
}

// Will set the component's ID.
func (e *Engine[Global]) AddDisabledComponent(entity *Entity, component Component) {
	components := e.GetComponents(component.GetKind().Id)

	component.EntityId = entity.Id
	components.Upsert(int(entity.Id), component)

	entity.markRemoved(component.GetKind().Id)
}

func (e *Engine[Global]) UpdateComponent(
	entity *Entity,
	kind *ComponentKind,
	f func(Component) Component,
) {
	components := e.GetComponents(kind.Id)
	id := int(entity.Id)

	if components.Has(id) {
		components.Update(id, f)
	}
}

func (e *Engine[Global]) OverwriteComponent(entity *Entity, component Component) {
	e.UpdateComponent(
		entity,
		component.GetKind(),
		func(_ Component) Component { return component },
	)
}

func (e *Engine[Global]) RemoveComponent(entity *Entity, kind *ComponentKind) {
	e.DisableComponent(entity, kind)

	components := e.GetComponents(kind.Id)

	id := entity.Id
	components.Remove(int(id))
}

func (e *Engine[Global]) EnableComponent(entity *Entity, kind *ComponentKind) {
	entity.addComponent(kind)

	// TODO: We can probably check less node kinds.
	for i := range len(NodeKinds) {
		nodeKind := &NodeKinds[i]
		componentKind := kind

		if componentKind.NodeValue%nodeKind.Value == 0 {
			if IsNodeKindSatisfied(nodeKind, *entity) {
				node := MakeStoredNode(nodeKind, entity.Id)
				nodes := e.GetStoredNodes(nodeKind.Id())
				nodes.Upsert(int(entity.Id), node)
			}
		}
	}
}

func (e *Engine[Global]) DisableComponent(entity *Entity, kind *ComponentKind) {
	happened := entity.removeComponent(kind)

	if !happened {
		return
	}

	// TODO: We can probably check less node kinds.
	for i := range len(NodeKinds) {
		nodeKind := &NodeKinds[i]
		componentKind := kind

		if componentKind.NodeValue%nodeKind.Value == 0 {
			if !IsNodeKindSatisfied(nodeKind, *entity) {
				nodes := e.GetStoredNodes(nodeKind.Id())
				nodes.Remove(int(entity.Id))
			}
		}
	}
}

// Enables the `current` component and disables every other component in
// `options` that is not `current`. Useful to implement state machines.
func (e *Engine[Global]) SwitchComponents(
	entity *Entity,
	options []*ComponentKind,
	current *ComponentKind,
) {
	e.EnableComponent(entity, current)

	for _, option := range options {
		if option.Id != current.Id {
			e.DisableComponent(entity, option)
		}
	}
}

func (e *Engine[Global]) AddSystem(system System[Global], priority int) {
	for _, eventKind := range system.EventKinds() {
		e.GetSystems(eventKind).Add(priority, system)
	}
}

func (e *Engine[Global]) RemoveSystem(system System[Global], priority int) {
	for _, eventKind := range system.EventKinds() {
		e.GetSystems(eventKind).RemoveP(priority, system)
	}
}

func (e *Engine[Global]) makeNode(sn StoredNode) Node {
	entity, ok := e.Entities.GetPtr(int(sn.EntityId))

	if !ok {
		panic("Could not find entity of node.")
	}

	node := MakeNode(entity)

	for _, componentKind := range sn.Kind.Components {
		components := e.GetComponents(componentKind.Id)
		component, ok := components.GetPtr(int(sn.EntityId))

		if !ok {
			panic("Could not find component of node.")
		}

		node.setComponent(componentKind, component)
	}

	return node
}

// BEWARE! The pointer is valid only until another operation is
// performed on components!
func (e *Engine[Global]) GetNodes(kind NodeKindId) []Node {
	nodes := []Node{}
	storedNodes := e.GetStoredNodes(kind).GetAllUnordered()

	for _, storedNode := range storedNodes {
		nodes = append(nodes, e.makeNode(storedNode))
	}

	return nodes
}

func (e *Engine[Global]) applyContext(ctx *EngineContext[Global]) {
	// First operations that do not change any array.

	for _, enableComponent := range ctx.enableComponents {
		e.EnableComponent(enableComponent.Entity, enableComponent.Kind)
	}

	for _, overwriteComponent := range ctx.overwriteComponents {
		e.OverwriteComponent(overwriteComponent.Entity, overwriteComponent.Component)
	}

	for _, disableComponent := range ctx.disableComponents {
		e.DisableComponent(disableComponent.Entity, disableComponent.Kind)
	}

	for _, switchComponent := range ctx.switchComponents {
		e.SwitchComponents(switchComponent.Entity, switchComponent.Options, switchComponent.Current)
	}

	// Invalidates component pointers.
	for _, removeComponent := range ctx.removeComponents {
		e.RemoveComponent(removeComponent.Entity, removeComponent.Kind)
	}

	// Does not use component pointers so it's fine.
	for _, addComponent := range ctx.addComponents {
		e.AddComponent(addComponent.Entity, addComponent.Component)
	}

	// Invalidates entity pointers.
	for _, removeEntity := range ctx.removeEntities {
		e.RemoveEntity(removeEntity)
	}

	// Does not use entity pointers so it's fine.
	for _, addEntity := range ctx.addEntities {
		e.AddEntity(addEntity)
	}

	// It's safe to modify queues basically any time.
	for _, addEvent := range ctx.addEvents {
		e.AddEvent(addEvent)
	}
}

type SystemToRun[Global any] struct {
	System System[Global]
	Ev     any
	Node   Node
}

func (e *Engine[Global]) dispatch(evKind EventKindId, ev any) {
	systems := e.GetSystems(evKind).GetAll()

	// TODO: Maybe have these at struct-level, just like the context, to
	// save on allocations.
	systemsToInit := make([]SystemToRun[Global], 0, systems.Len())
	systemsToRun := make([]SystemToRun[Global], 0, systems.Len())

	for hasPrio := !systems.IsEmpty(); hasPrio; hasPrio = systems.NextPriority() {
		e.Context.Clean()
		systemsToInit = systemsToInit[:0]
		systemsToRun = systemsToRun[:0]

		for system := systems.NextInPriority(); system != nil; system = systems.NextInPriority() {
			nodeKind := system.NodeKind()
			storedNodes := e.GetStoredNodes(nodeKind.Id()).GetAllUnordered()

			for i, storedNode := range storedNodes {
				node := e.makeNode(storedNode)

				hasSeen := storedNode.See(system.Value)
				if !hasSeen {
					systemsToInit = append(systemsToInit, SystemToRun[Global]{
						System: *system,
						Ev:     ev,
						Node:   node,
					})

					// Save the `See`.
					storedNodes[i] = storedNode
				}

				systemsToRun = append(systemsToRun, SystemToRun[Global]{
					System: *system,
					Ev:     ev,
					Node:   node,
				})
			}
		}

		// TODO: This can probably be parallelized just fine with some
		// syncing in EngineContext and Global.
		for _, systemInit := range systemsToInit {
			systemInit.System.Init(&e.Context, systemInit.Ev, systemInit.Node)
		}

		e.applyContext(&e.Context)

		for _, systemRun := range systemsToRun {
			systemRun.System.Handle(&e.Context, systemRun.Ev, systemRun.Node)
		}

		e.applyContext(&e.Context)
	}
}

type EngineFrameContext = FrameEvent
type EngineRenderContext = RenderEvent
type EngineResizeContext = ResizeEvent

func (e *Engine[Global]) Tick(data EngineFrameContext) {
	for ev, ok := e.Events.Dequeue(); ok; ev, ok = e.Events.Dequeue() {
		e.dispatch(ev.KindId, ev.Data)
	}

	e.dispatch(FrameEventKindId, data)
}

func (e *Engine[Global]) Render(data EngineRenderContext) {
	e.dispatch(RenderEventKindId, data)
}

func (e *Engine[Global]) AddEvent(ev Event) {
	if ev.KindId == RenderEventKindId {
		panic("Tried to enqueue Render event on the engine. That will surely not work.")
	}

	e.Events.Enqueue(ev)
}
