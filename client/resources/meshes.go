package resources

import (
	"unsafe"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// There could be different kinds of meshes at some point.

type Mesh struct {
	// TODO: Do we really need to store all of this here?
	Coords    []digimath.Vec3
	Indices   []uint32
	TexCoords []digimath.Vec2

	GLVertexArrayId  uint32
	GLVertexBufferId uint32
}

func MakeMesh(
	coords []digimath.Vec3,
	texCoords []digimath.Vec2,
	indices []uint32,
) Mesh {
	mesh := Mesh{}

	mesh.Coords = coords
	mesh.TexCoords = texCoords
	mesh.Indices = indices

	coordsSize := int(unsafe.Sizeof(coords[0]))*len(coords)
	texCoordsSize := int(unsafe.Sizeof(texCoords[0]))*len(texCoords)

	totalBufferSize := coordsSize + texCoordsSize

	gl.GenBuffers(1, &mesh.GLVertexBufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.GLVertexBufferId)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		totalBufferSize,
		nil,
		gl.STATIC_DRAW,
	)

	gl.BufferSubData(
		gl.ARRAY_BUFFER,
		0,
		coordsSize,
		gl.Ptr(coords[:]),
	)

	gl.BufferSubData(
		gl.ARRAY_BUFFER,
		coordsSize,
		texCoordsSize,
		gl.Ptr(texCoords[:]),
	)

	gl.GenVertexArrays(1, &mesh.GLVertexArrayId)
	gl.BindVertexArray(mesh.GLVertexArrayId)
	defer gl.BindVertexArray(0)

	// Coords
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		0,
		uintptr(0),
	)

	// TexCoords
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(
		1,
		2,
		gl.FLOAT,
		false,
		0,
		uintptr(coordsSize),
	)

	indexSize := int(unsafe.Sizeof(indices[0]))*len(indices)

	var indexBufferId uint32
	gl.GenBuffers(1, &indexBufferId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBufferId)

	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		indexSize,
		gl.Ptr(indices[:]),
		gl.STATIC_DRAW,
	)

	return mesh
}

func (m *Mesh) IndexCount() int32 {
	return int32(len(m.Indices))
}

func MakeQuadMesh(
	rect digimath.Rect,
) Mesh {
	coords := [4]digimath.Vec3{
		rect.GetBottomRight(),
		rect.GetTopRight(),
		rect.GetTopLeft(),
		rect.GetBottomLeft(),
	}

	texCoords := [4]digimath.Vec2{
		digimath.MakeVec2(1, 1),
		digimath.MakeVec2(1, 0),
		digimath.MakeVec2(0, 0),
		digimath.MakeVec2(0, 1),
	}

	indices := [6]uint32{
		0, 1, 2, 2, 3, 0,
	}

	return MakeMesh(coords[:], texCoords[:], indices[:])
}

func (m *Mesh) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, m.GLVertexBufferId)
	gl.BindVertexArray(m.GLVertexArrayId)
}

func (m *Mesh) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}
