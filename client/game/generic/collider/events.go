package collider

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
)

// Collision masks (not implemented yet) may make it so that one thing
// collides with another, but that one does not collide back. If that's
// not the case, two events will be generated on each collision: One
// with each being active and the other passive.

var CollisionOnEventKindId = engine.MakeEventKindId()
var CollisionOffEventKindId = engine.MakeEventKindId()

type CollisionOnEvent struct {
	Active  engine.EntityId
	Passive engine.EntityId
	Point   digimath.Vec3
}

type CollisionOffEvent struct {
	Active  engine.EntityId
	Passive engine.EntityId
	Point   digimath.Vec3
}

// TODO: Add collision with quad (not plane)
// TODO: Add collision to camera (ray)
// TODO: Track the closest thing the camera is pointing at
