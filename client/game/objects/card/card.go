package card

import "github.com/CrimsonSarah/cto/client/digimath"

type CardVertexData struct {
	Coords    [4]digimath.Vec3
	TexCoords [4]digimath.Vec2
}

var CardVertices = CardVertexData{
	Coords: [4]digimath.Vec3{
		digimath.MakeVec3(-0.358, -0.5, 0),
		digimath.MakeVec3(-0.358, 0.5, 0),
		digimath.MakeVec3(0.358, 0.5, 0),
		digimath.MakeVec3(0.358, -0.5, 0),
	},
	TexCoords: [4]digimath.Vec2{
		digimath.MakeVec2(0, 1),
		digimath.MakeVec2(0, 0),
		digimath.MakeVec2(1, 0),
		digimath.MakeVec2(1, 1),
	},
}

var CardVertexIndices = [6]uint32{
	0, 1, 2, 2, 3, 0,
}

type Card struct {
	Code string
	Name string
}

func MakeCard(code, name string) Card {
	return Card{
		Name: name,
		Code: code,
	}
}

// See `world/object.go`.
func (o Card) Intersects(p digimath.Vec2) bool {
	bottomLeft := CardVertices.Coords[0]
	topRight := CardVertices.Coords[2]

	return p.X() >= bottomLeft.X() && p.X() <= topRight.X() &&
		p.Y() >= bottomLeft.Y() && p.Y() <= topRight.Y()
}
