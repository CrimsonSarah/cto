package camera

import (
	"log"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/collider"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
)

func AddCameraSystems(
	e *engine.Engine[global.GameGlobals],
) {
	e.AddSystem(CameraInputSystem, 0)
	e.AddSystem(CameraDefaultSystem, 1)
	e.AddSystem(CameraExploreSystem, 1)
}

func AddCamera(
	e *engine.Engine[global.GameGlobals],
	windowWidth, windowHeight int,
) (*engine.Entity, error) {
	camera := MakeCameraCameraComponent(windowWidth, windowHeight)
	input := MakeCameraInputComponent()
	transform := transform.MakeTransformComponent()

	collider := collider.MakeColliderComponentRay(
		digimath.Vec3Zero,
		digimath.MakeRay(
			digimath.Vec3Zero, digimath.MakeVec3(0, 0, 1),
		),
	)

	stateDefault := MakeCameraStateDefaultComponent()
	stateExplore := MakeCameraStateExploreComponent()

	ent := e.AddNewEntity()
	e.AddComponent(ent, camera)
	e.AddComponent(ent, input)
	e.AddComponent(ent, transform)
	e.AddComponent(ent, collider)

	e.AddComponent(ent, stateDefault)
	e.AddDisabledComponent(ent, stateExplore)

	log.Printf("Created camera entity with ID %v\n", ent.Id)
	return ent, nil
}
