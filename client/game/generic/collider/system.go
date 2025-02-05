package collider

import (
	"log"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
)

type ColliderSystemImpl struct{}

var ColliderSystem = engine.MakeSystem[global.GameGlobals](
	"Collider",
	ColliderSystemImpl{},
)

func (s ColliderSystemImpl) EventKinds() []engine.EventKindId {
	return []engine.EventKindId{engine.RenderEventKindId}
}

func (s ColliderSystemImpl) NodeKind() *engine.NodeKind {
	return &ColliderNodeKind
}

func (s ColliderSystemImpl) Init(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
}

func (s ColliderSystemImpl) Handle(
	ctx *engine.EngineContext[global.GameGlobals],
	event any,
	node engine.Node,
) {
	otherNodes := ctx.GetNodes(ColliderNodeKind.Id())

	for _, otherNode := range otherNodes {
		if otherNode.Entity.Id == node.Entity.Id {
			continue
		}

		result := Intersects(node, otherNode)

		if result.DidCollide {
			log.Printf("collision of %d and %d @ %v\n",
				result.Active,
				result.Passive,
				result.Point,
			)

			ctx.AddEvent(engine.Event{
				KindId: CollisionOnEventKindId,
				Data: CollisionOnEvent{
					Active:  result.Active,
					Passive: result.Passive,
					Point:   result.Point,
				},
			})
		}
	}
}

type IntersectResult struct {
	DidCollide bool
	Point      digimath.Vec3
	Active     engine.EntityId
	Passive    engine.EntityId
}

func Intersects(n1, n2 engine.Node) IntersectResult {
	c1 := n1.GetData(&ColliderComponentKind).(Collider)
	t1 := n1.GetData(&transform.TransformComponentKind).(transform.Transform)

	c2 := n2.GetData(&ColliderComponentKind).(Collider)
	t2 := n2.GetData(&transform.TransformComponentKind).(transform.Transform)

	didCollide := false
	point := digimath.Vec3Zero

	active := n1.Entity.Id
	passive := n2.Entity.Id

	switch shape1 := c1.Shape.(type) {
	case digimath.Ray:
		shape1 = t1.ApplyToRay(shape1, c1.Pivot)

		switch shape2 := c2.Shape.(type) {
		case digimath.Plane:
			shape2 = t2.ApplyToPlane(shape2, c2.Pivot)
			didCollide, point = digimath.IntersectRayPlane(shape1, shape2)
		case digimath.Rect:
			shape2 = t2.ApplyToRect(shape2, c2.Pivot)
			didCollide, point = digimath.IntersectRayRect(shape1, shape2)
		}
	case digimath.Plane:
		shape1 = t1.ApplyToPlane(shape1, c1.Pivot)

		switch shape2 := c2.Shape.(type) {
		case digimath.Ray:
			shape2 = t2.ApplyToRay(shape2, c2.Pivot)
			didCollide, point = digimath.IntersectRayPlane(shape2, shape1)
		}
	case digimath.Rect:
		shape1 = t1.ApplyToRect(shape1, c1.Pivot)

		switch shape2 := c2.Shape.(type) {
		case digimath.Ray:
			shape2 = t2.ApplyToRay(shape2, c2.Pivot)
			didCollide, point = digimath.IntersectRayRect(shape2, shape1)
		}
	}

	return IntersectResult{
		DidCollide: didCollide,
		Point:      point,
		Active:     active,
		Passive:    passive,
	}
}
