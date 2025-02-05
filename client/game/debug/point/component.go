package point

import (
	"fmt"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/resources"
)

var DebugPointComponentKind = engine.MakeComponentKind("DebugPoint")

type DebugPoint struct {
	Color         digimath.Vec3
	Mesh          resources.Mesh
	ShaderProgram resources.ShaderProgram

	ProjectionUniformLocation int32
	CameraUniformLocation     int32
	TransformUniformLocation  int32
}

func MakeDebugPoint(color digimath.Vec3) (DebugPoint, error) {
	p := DebugPoint{
		Color: color,
	}

	vertShader, err :=
		resources.MakeShader(
			resources.ResPath(
				"resources/shaders/debug/point/vert.glsl",
			),
			resources.VertShaderType,
		)

	if err != nil {
		return p, err
	}

	fragShader, err :=
		resources.MakeShader(
			resources.ResPath(
				"resources/shaders/debug/point/frag.glsl",
			),
			resources.FragShaderType,
		)

	if err != nil {
		return p, err
	}

	rect := digimath.MakeProfileRect(0.1, 0.1)
	p.Mesh = resources.MakeQuadMesh(rect)

	p.ShaderProgram = resources.MakeShaderProgram()
	p.ShaderProgram.AttachShader(vertShader)
	p.ShaderProgram.AttachShader(fragShader)
	p.ShaderProgram.Build()

	p.ProjectionUniformLocation =
		p.ShaderProgram.GetUniformLocation("u_Projection")
	p.CameraUniformLocation =
		p.ShaderProgram.GetUniformLocation("u_Camera")
	p.TransformUniformLocation =
		p.ShaderProgram.GetUniformLocation("u_Transform")

	return p, nil
}

func (d *DebugPoint) Bind() {
	d.Mesh.Bind()
	d.ShaderProgram.Bind()
}

func (p DebugPoint) GetKind() *engine.ComponentKind {
	return &DebugPointComponentKind
}

func MakeDebugPointComponent(
	color digimath.Vec3,
) engine.Component {
	p, err := MakeDebugPoint(color)

	if err != nil {
		panic(fmt.Errorf("Could not create debug point: %w", err))
	}

	return engine.MakeComponent(p)
}
