package debug

import (
	"log"
	"unsafe"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/objects/debug"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/resources"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Doesn't actually add anything.
type RenderableDebugPoint struct {
	*world.Placed[debug.DebugPoint]
}

type DebugPointRenderer struct {
	World *world.World

	VertexArrayId  uint32
	VertexBufferId uint32
	ProgramId      uint32

	ProjectionUniformLocation       int32
	TransformUniformLocation        int32
	ColorUniformLocation            int32
	CenterUniformLocation           int32
	RadiusUniformLocation           int32
	WindowDimensionsUniformLocation int32
}

const RADIUS float32 = 0.025

func (r *DebugPointRenderer) Init(world *world.World) {
	r.World = world

	gl.GenVertexArrays(1, &r.VertexArrayId)
	gl.BindVertexArray(r.VertexArrayId)
	defer gl.BindVertexArray(0)

	gl.GenBuffers(1, &r.VertexBufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.VertexBufferId)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		int(unsafe.Sizeof(debug.DebugPointVertexData{})),
		nil,
		gl.DYNAMIC_DRAW,
	)

	// Coords.
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		0,
		unsafe.Offsetof(debug.DebugPointVertexData{}.Coords),
	)

	var indexBufferId uint32
	gl.GenBuffers(1, &indexBufferId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBufferId)

	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		int(unsafe.Sizeof(debug.DebugPointVertexIndices)),
		gl.Ptr(debug.DebugPointVertexIndices[:]),
		gl.STATIC_DRAW,
	)

	// Setup shaders.

	r.ProgramId = gl.CreateProgram()

	r.attachShader(
		gl.VERTEX_SHADER,
		"resources/shaders/debug/point/vert.glsl",
	)

	r.attachShader(
		gl.FRAGMENT_SHADER,
		"resources/shaders/debug/point/frag.glsl",
	)

	gl.LinkProgram(r.ProgramId)
	gl.ValidateProgram(r.ProgramId)
	gl.UseProgram(r.ProgramId)

	r.ProjectionUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_Projection\000"),
	)

	gl.UniformMatrix4fv(
		r.ProjectionUniformLocation,
		1,
		false,
		&world.Projection[0],
	)

	r.TransformUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_Transform\000"),
	)

	r.ColorUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_Color\000"),
	)

	r.CenterUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_Center\000"),
	)

	r.RadiusUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_Radius\000"),
	)

	gl.Uniform1f(r.RadiusUniformLocation, RADIUS)

	r.WindowDimensionsUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_WindowDimensions\000"),
	)

	gl.Uniform2f(
		r.WindowDimensionsUniformLocation,
		float32(world.WindowWidth),
		float32(world.WindowHeight),
	)
}

func (r *DebugPointRenderer) Configure() {
	gl.UseProgram(r.ProgramId)
	gl.UniformMatrix4fv(
		r.ProjectionUniformLocation,
		1,
		false,
		&r.World.Projection[0],
	)
	gl.Uniform2f(
		r.WindowDimensionsUniformLocation,
		float32(r.World.WindowWidth),
		float32(r.World.WindowHeight),
	)
}

func (r *DebugPointRenderer) attachShader(
	xtype uint32,
	path resources.ResPath,
) {
	shader, err := resources.ReadShader(path)

	if err != nil {
		log.Fatalln("Could not load vertex shader for debug point", err)
	}

	shaderId := gl.CreateShader(xtype)
	shaderCompat, free := gl.Strs(shader)

	shaderLength := int32(len(shader))
	gl.ShaderSource(shaderId, 1, shaderCompat, &shaderLength)
	gl.CompileShader(shaderId)
	free()

	gl.AttachShader(r.ProgramId, shaderId)
	gl.DeleteShader(shaderId)
}

func (r *DebugPointRenderer) MakeRenderableDebugPoint(
	point *world.Placed[debug.DebugPoint],
) RenderableDebugPoint {
	return RenderableDebugPoint{
		point,
	}
}

func (r *DebugPointRenderer) RenderDebugPoint(
	o *RenderableDebugPoint,
) {
	gl.BindBuffer(gl.ARRAY_BUFFER, r.VertexBufferId)
	gl.BindVertexArray(r.VertexArrayId)
	defer gl.BindVertexArray(0)

	gl.UseProgram(r.ProgramId)

	// TODO: TEST
	RADIUS := float32(1)

	center := o.Obj.Center
	bottomLeft := center.Add(digimath.MakeVec3(-RADIUS, -RADIUS, 0))
	topLeft := center.Add(digimath.MakeVec3(-RADIUS, RADIUS, 0))
	topRight := center.Add(digimath.MakeVec3(RADIUS, RADIUS, 0))
	bottomRight := center.Add(digimath.MakeVec3(RADIUS, -RADIUS, 0))

	buffer := debug.DebugPointVertexData{
		Coords: [4]digimath.Vec3{
			bottomLeft,
			topLeft,
			topRight,
			bottomRight,
		},
	}

	gl.BufferSubData(
		gl.ARRAY_BUFFER,
		0,
		int(unsafe.Sizeof(debug.DebugPointVertexData{})),
		gl.Ptr(buffer.Coords[:]),
	)

	transform := o.Transform.ToMatrix()
	gl.UniformMatrix4fv(
		r.TransformUniformLocation,
		1,
		false,
		&transform[0],
	)

	gl.Uniform3f(
		r.ColorUniformLocation,
		o.Obj.Color[0],
		o.Obj.Color[1],
		o.Obj.Color[2],
	)

	gl.Uniform2f(
		r.CenterUniformLocation,
		center[0],
		center[1],
	)

	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(debug.DebugPointVertexIndices)),
		gl.UNSIGNED_INT,
		nil,
	)
}
