package engine

var componentKindIdCounter = 0

type ComponentKindId int

// A component kind is used to create nodes from components. Every
// component shall have a component kind.
type ComponentKind struct {
	Id   ComponentKindId
	Name string

	// See `node.go`.
	NodeValue uint64
}

func MakeComponentKind(name string) ComponentKind {
	id := componentKindIdCounter
	componentKindIdCounter += 1

	return ComponentKind{
		Id:        ComponentKindId(id),
		Name:      name,
		NodeValue: 1,
	}
}

// This should be called by the node kind itself on creation.
func (c *ComponentKind) addNode(node uint64) {
	c.NodeValue *= node
}

type ComponentData interface {
	GetKind() *ComponentKind
}

type Component struct {
	Data ComponentData

	// Inserted after creation.
	EntityId EntityId
}

func MakeComponent(data ComponentData) Component {
	return Component{
		Data: data,
	}
}

func (c Component) GetKind() *ComponentKind {
	return c.Data.GetKind()
}

func (c Component) Update(data ComponentData) Component {
	c.Data = data
	return c
}

// IMPORTANT: Components are only truly equal if they are both of the
// same kind and of the same entity, but this does NOT compare kinds.
// That's not a problem when traversing components by kind, but will
// be a problem if traversing components by entity.
func (c1 Component) Equals(c2 Component) bool {
	return c1.EntityId == c2.EntityId
}
