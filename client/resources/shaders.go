package resources

import "github.com/go-gl/gl/v3.3-core/gl"

type Shader struct {
	GLShaderId uint32
}

type ShaderType uint32

const (
	VertShaderType = ShaderType(gl.VERTEX_SHADER)
	FragShaderType = ShaderType(gl.FRAGMENT_SHADER)
)

func MakeShader(path ResPath, xtype ShaderType) (Shader, error) {
	shader := Shader{}
	contentsB, err := ReadNullTerminated(path)

	if err != nil {
		return Shader{}, err
	}

	contents := string(contentsB)
	contentsCompat, free := gl.Strs(contents)
	contentsLen := int32(len(contents))

	shader.GLShaderId = gl.CreateShader(uint32(xtype))
	gl.ShaderSource(shader.GLShaderId, 1, contentsCompat, &contentsLen)
	gl.CompileShader(shader.GLShaderId)
	free()

	return shader, nil
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.GLShaderId)
}
