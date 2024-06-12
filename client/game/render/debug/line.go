package debug

import (
	"log"
	"unsafe"

	"github.com/CrimsonSarah/cto/client/game/objects/debug"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/resources"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Doesn't actually add anything.
type RenderableDebugLines struct {
	Lines *world.Placed[debug.DebugLines]
}

type DebugLineRenderer struct {
	World *world.World

	VertexArrayId  uint32
	VertexBufferId uint32
	ProgramId      uint32

	ProjectionUniformLocation int32
	TransformUniformLocation  int32
	ColorUniformLocation      int32

	CurrentLineCount int
}

func (r *DebugLineRenderer) Init(world *world.World) {
	r.World = world
	r.CurrentLineCount = 0

	gl.GenBuffers(1, &r.VertexBufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.VertexBufferId)

	gl.BufferData(
		gl.ARRAY_BUFFER,

		// Fancy way of saying zero.
		r.CurrentLineCount*
			int(unsafe.Sizeof(debug.DebugLineVertexData{})),

		nil,
		gl.DYNAMIC_DRAW,
	)

	gl.GenVertexArrays(1, &r.VertexArrayId)
	gl.BindVertexArray(r.VertexArrayId)
	defer gl.BindVertexArray(0)

	// Coords.
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		int32(unsafe.Sizeof(debug.DebugLineVertexData{})),
		unsafe.Offsetof(debug.DebugLineVertexData{}.Coords),
	)

	// Setup shaders.

	r.ProgramId = gl.CreateProgram()

	r.attachShader(
		gl.VERTEX_SHADER,
		"resources/shaders/debug/line/vert.glsl",
	)

	r.attachShader(
		gl.FRAGMENT_SHADER,
		"resources/shaders/debug/line/frag.glsl",
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
}

func (r *DebugLineRenderer) Configure() {
	gl.UseProgram(r.ProgramId)
	gl.UniformMatrix4fv(
		r.ProjectionUniformLocation,
		1,
		false,
		&r.World.Projection[0],
	)
}

func (r *DebugLineRenderer) attachShader(
	xtype uint32,
	path resources.ResPath,
) {
	shader, err := resources.ReadShader(path)

	if err != nil {
		log.Fatalln("Could not load vertex shader for debug line", err)
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

func (r *DebugLineRenderer) MakeRenderableDebugLines(
	lines *world.Placed[debug.DebugLines],
) RenderableDebugLines {
	return RenderableDebugLines{
		Lines: lines,
	}
}

// Depends on the layout of DebugLine to work.
func (r *DebugLineRenderer) RenderDebugLine(
	o *RenderableDebugLines,
) {
	count := len(o.Lines.Obj.Lines)

	if count <= 0 {
		return
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, r.VertexBufferId)
	gl.BindVertexArray(r.VertexArrayId)
	defer gl.BindVertexArray(0)

	gl.UseProgram(r.ProgramId)

	if count > r.CurrentLineCount {
		gl.BufferData(
			gl.ARRAY_BUFFER,
			2*count*
				int(unsafe.Sizeof(debug.
					DebugLineVertexData{})),
			nil,
			gl.DYNAMIC_DRAW,
		)
	}

	r.CurrentLineCount = count

	gl.BufferSubData(
		gl.ARRAY_BUFFER,
		0,
		2*count*
			int(unsafe.Sizeof(debug.DebugLineVertexData{})),
		gl.Ptr(o.Lines.Obj.Lines[:]),
	)

	transform := o.Lines.Transform.ToMatrix()
	gl.UniformMatrix4fv(
		r.TransformUniformLocation,
		1,
		false,
		&transform[0],
	)

	gl.Uniform3f(
		r.ColorUniformLocation,
		o.Lines.Obj.Color[0],
		o.Lines.Obj.Color[1],
		o.Lines.Obj.Color[2],
	)

	gl.DrawArrays(gl.LINES, 0, int32(count*2))
}
