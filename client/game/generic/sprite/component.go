package sprite

import (
	"github.com/CrimsonSarah/cto/client/digigl"
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/resources"
)

var SpriteComponentKind = engine.MakeComponentKind("Sprite")

type Sprite struct {
	Texture       resources.BoundTexture
	Mesh          resources.Mesh
	ShaderProgram resources.ShaderProgram

	ProjectionUniformLocation int32
	CameraUniformLocation     int32
	TransformUniformLocation  int32
}

// This will create a rect with the proportions of the loaded texture
// and longest side of length 1.
func MakeDefaultSpriteRect(tex resources.Texture) digimath.Rect {
	bounds := tex.Bounds()
	absWidth := bounds.Max.X
	absHeight := bounds.Max.Y

	maxLength := max(absWidth, absHeight)
	relWidth := float32(absWidth) / float32(maxLength)
	relHeight := float32(absHeight) / float32(maxLength)

	return digimath.MakeProfileRect(relWidth, relHeight)
}

func MakeSprite(
	rect digimath.Rect,
	tex resources.Texture,
	shaders ...resources.Shader,
) Sprite {
	s := Sprite{
		Texture: resources.MakeBoundTexture(tex, digigl.SpriteTextureUnit),
	}

	s.Mesh = resources.MakeQuadMesh(rect)
	s.ShaderProgram = resources.MakeShaderProgram()

	for _, shader := range shaders {
		s.ShaderProgram.AttachShader(shader)
	}

	s.ShaderProgram.Build()

	s.ProjectionUniformLocation =
		s.ShaderProgram.GetUniformLocation("u_Projection")
	s.CameraUniformLocation =
		s.ShaderProgram.GetUniformLocation("u_Camera")
	s.TransformUniformLocation =
		s.ShaderProgram.GetUniformLocation("u_Transform")

	return s
}

func (s *Sprite) Bind() {
	s.Mesh.Bind()
	s.ShaderProgram.Bind()
	s.Texture.Bind()
}

func (s Sprite) GetKind() *engine.ComponentKind {
	return &SpriteComponentKind
}

func MakeSpriteComponent(
	rect digimath.Rect,
	tex resources.Texture,
	shaders ...resources.Shader,
) engine.Component {
	return engine.MakeComponent(MakeSprite(rect, tex, shaders...))
}
