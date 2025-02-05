package engine

// Each node shall have an associated prime number. ComponentKinds will
// contain the multiplication of the prime numbers of each node it can
// make, and that will be used to build the nodes automatically as
// needed.
type NodeKind struct {
	Name       string
	Value      uint64
	Components []*ComponentKind
}

var NodeKinds []NodeKind
var nodeKindValueCounter = 0

func MakeNodeKind(name string, components ...*ComponentKind) NodeKind {
	counter := nodeKindValueCounter
	nodeKindValueCounter += 1

	value := Primes[counter]

	for _, component := range components {
		component.addNode(value)
	}

	kind := NodeKind{
		Name:       name,
		Value:      value,
		Components: components,
	}

	NodeKinds = append(NodeKinds, kind)
	return kind
}

type NodeKindId uint64

func (n NodeKind) Id() NodeKindId {
	return NodeKindId(n.Value)
}

// A StoredNode exists with a given Entity if that entity is capable of
// producing nodes of kind Kind.
type StoredNode struct {
	EntityId EntityId
	Kind     *NodeKind

	// The multiplication of the values of every system that has been run
	// on this node. If a system will run and this is not a multiple of
	// its value, it shall be Init first.
	SeenSystems uint64
}

func MakeStoredNode(kind *NodeKind, entity EntityId) StoredNode {
	return StoredNode{
		EntityId:    entity,
		Kind:        kind,
		SeenSystems: 1,
	}
}

func (n StoredNode) HasSeen(systemValue uint64) bool {
	// This is quite a price to pay. It should only really pay off if
	// there are many systems.
	return n.SeenSystems % systemValue == 0
}

// Returns `true` if it had been seen before, `false` otherwise.
func (n *StoredNode) See(systemValue uint64) bool {
	if n.HasSeen(systemValue)	{
		return true
	}

	n.SeenSystems *= systemValue
	return false
}

// IMPORTANT: StoredNodes are only truly equal if they are both of the
// same kind and of the same entity, but this does NOT compare kinds.
// That's not a problem when traversing nodes by kind, but will
// be a problem if traversing nodes by entity.
func (n1 StoredNode) Equals(n2 StoredNode) bool {
	return n1.EntityId == n2.EntityId
}

// Returns whether it is possible to build a node of the given kind
// using the components in the given entity. We use the `nodeValue`,
// which is the multiplication of the node values of its components.
// Since nodes are always composed of distinct components, we check if
// the entity's node value is divisible by $ n * kind.Value $, where
// $ n $ is the amount of components needed to construct the node.
// That will only be the case if there are $ n $ distinct components
// that make up that node, which means there are all of them, so it is
// possible to build the node.
func IsNodeKindSatisfied(kind *NodeKind, entity Entity) bool {
	componentCount := len(kind.Components)
	value := entity.NodeValue
	result := uint64(0)

	// TODO: Just exponentiate and check once..?
	for range componentCount {
		result |= value % kind.Value
		value = value / kind.Value
	}

	return result == 0
}

func (n *StoredNode) SetEntity(id EntityId) {
	n.EntityId = id
}

func (n StoredNode) GetEntity() EntityId {
	return n.EntityId
}

type Node struct {
	Entity     *Entity
	components map[ComponentKindId]*Component
}

func MakeNode(entity *Entity) Node {
	return Node{
		Entity:     entity,
		components: make(map[ComponentKindId]*Component),
	}
}

func (n *Node) setComponent(kind *ComponentKind, component *Component) {
	n.components[kind.Id] = component
}

func (n *Node) GetComponent(kind *ComponentKind) Component {
	return *n.components[kind.Id]
}

func (n *Node) GetData(kind *ComponentKind) any {
	return n.components[kind.Id].Data
}
