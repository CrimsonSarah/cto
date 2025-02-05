package engine

// 0 is reserved.
var entityIdCounter = 1

type EntityId int

type Entity struct {
	Id        EntityId
	NodeValue uint64

	// Used to prevent removing the same component twice.
	DisabledComponents map[ComponentKindId]bool
}

func MakeEntity() Entity {
	id := entityIdCounter
	entityIdCounter += 1

	return Entity{
		Id:         EntityId(id),
		NodeValue:  1,

		DisabledComponents: make(map[ComponentKindId]bool),
	}
}

func (e *Entity) addComponent(kind *ComponentKind) {
	delete(e.DisabledComponents, kind.Id)
	e.NodeValue *= kind.NodeValue
}

// Return of `true` indicate removal happened. Return of `false` means
// nothing happened because it was already removed.
func (e *Entity) removeComponent(kind *ComponentKind) bool {
	if e.isRemoved(kind.Id) {
		return false
	}

	e.DisabledComponents[kind.Id] = true
	e.NodeValue /= kind.NodeValue
	return true
}

func (e *Entity) isRemoved(id ComponentKindId) bool {
	value, ok := e.DisabledComponents[id]
	return value && ok
}

func (e *Entity) markRemoved(id ComponentKindId) {
	e.DisabledComponents[id] = true
}

func (a Entity) Equals(b Entity) bool {
	return a.Id == b.Id
}
