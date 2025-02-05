package resources

import "github.com/go-gl/gl/v3.3-core/gl"

type ShaderProgram struct {
	GLProgramId uint32
}

func MakeShaderProgram() ShaderProgram {
	program := ShaderProgram{}
	program.GLProgramId = gl.CreateProgram()

	return program
}

func (p *ShaderProgram) AttachShader(shader Shader) {
	gl.AttachShader(p.GLProgramId, shader.GLShaderId)
}

func (p *ShaderProgram) Build() {
	gl.LinkProgram(p.GLProgramId)
	gl.ValidateProgram(p.GLProgramId)
	gl.UseProgram(p.GLProgramId)
}

func (p *ShaderProgram) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(
		p.GLProgramId,
		gl.Str(name+"\000"),
	)
}

func (p *ShaderProgram) Bind() {
	gl.UseProgram(p.GLProgramId)
}
