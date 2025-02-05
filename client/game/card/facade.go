package card

import (
	"log"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/collider"
	"github.com/CrimsonSarah/cto/client/game/generic/sprite"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/resources"
)

func AddCard[Global any](
	e *engine.Engine[Global],
) (*engine.Entity, error) {
	var ent *engine.Entity = nil

	tex, err := resources.MakeTexture(
		resources.ResPath("resources/textures/cards/BT5-103.jpg"),
	)

	if err != nil {
		return ent, err
	}

	vert, err := resources.MakeShader(
		resources.ResPath("resources/shaders/cards/vert.glsl"),
		resources.VertShaderType,
	)

	if err != nil {
		return ent, err
	}

	frag, err := resources.MakeShader(
		resources.ResPath("resources/shaders/cards/frag.glsl"),
		resources.FragShaderType,
	)

	if err != nil {
		return ent, err
	}

	rect := sprite.MakeDefaultSpriteRect(tex)

	transform := transform.MakeTransformComponent()
	sprite := sprite.MakeSpriteComponent(rect, tex, vert, frag)
	collider := collider.MakeColliderComponentRect(digimath.Vec3Zero, rect)

	ent = e.AddNewEntity()
	e.AddComponent(ent, transform)
	e.AddComponent(ent, sprite)
	e.AddComponent(ent, collider)

	log.Printf("Created card entity with ID %v\n", ent.Id)
	return ent, nil
}
