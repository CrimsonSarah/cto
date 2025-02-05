package collider

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
)

var ColliderComponentKind = engine.MakeComponentKind("Collider")

type Collider struct {
	Pivot digimath.Vec3
	Shape any
}

func MakeColliderRay(pivot digimath.Vec3, ray digimath.Ray) Collider {
	return Collider{
		Pivot: pivot,
		Shape: ray,
	}
}

func MakeColliderPlane(pivot digimath.Vec3, plane digimath.Plane) Collider {
	return Collider{
		Pivot: pivot,
		Shape: plane,
	}
}

func MakeColliderRect(pivot digimath.Vec3, rect digimath.Rect) Collider {
	return Collider{
		Pivot: pivot,
		Shape: rect,
	}
}

func (c Collider) GetKind() *engine.ComponentKind {
	return &ColliderComponentKind
}

func MakeColliderComponentRay(pivot digimath.Vec3, r digimath.Ray) engine.Component {
	return engine.MakeComponent(MakeColliderRay(pivot, r))
}

func MakeColliderComponentPlane(pivot digimath.Vec3, p digimath.Plane) engine.Component {
	return engine.MakeComponent(MakeColliderPlane(pivot, p))
}

func MakeColliderComponentRect(pivot digimath.Vec3, r digimath.Rect) engine.Component {
	return engine.MakeComponent(MakeColliderRect(pivot, r))
}
