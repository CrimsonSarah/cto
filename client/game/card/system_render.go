package card

// import (
// 	"fmt"
// 	"log"
// 	"unsafe"
// 
// 	"github.com/CrimsonSarah/cto/client/digigl"
// 	"github.com/CrimsonSarah/cto/client/engine"
// 	"github.com/CrimsonSarah/cto/client/game/components"
// 	"github.com/CrimsonSarah/cto/client/game/objects/card"
// 	"github.com/CrimsonSarah/cto/client/game/world"
// 	"github.com/CrimsonSarah/cto/client/resources"
// 	"github.com/go-gl/gl/v3.3-core/gl"
// )
// 
// type CardRendererSystem struct {
// 	World *world.World
// 
// 	VertexArrayId  uint32
// 	VertexBufferId uint32
// 	ProgramId      uint32
// 
// 	ProjectionUniformLocation int32
// 	CameraUniformLocation     int32
// 	TransformUniformLocation  int32
// 
// 	CardTextures map[string]uint32
// }
// 
// func (s *CardRendererSystem) Init(c engine.InitContext) {
// 	s.World = c.World
// 
// 	// Setup vertex and index buffers
// 	log.Println("Initialiazing cards")
// 
// 	gl.GenBuffers(1, &s.VertexBufferId)
// 	gl.BindBuffer(gl.ARRAY_BUFFER, s.VertexBufferId)
// 
// 	gl.BufferData(
// 		gl.ARRAY_BUFFER,
// 		int(unsafe.Sizeof(card.CardVertices)),
// 		gl.Ptr(card.CardVertices.Coords[:]),
// 		gl.STATIC_DRAW,
// 	)
// 
// 	gl.GenVertexArrays(1, &s.VertexArrayId)
// 	gl.BindVertexArray(s.VertexArrayId)
// 	defer gl.BindVertexArray(0)
// 
// 	gl.EnableVertexAttribArray(0)
// 	gl.VertexAttribPointerWithOffset(
// 		0,
// 		3,
// 		gl.FLOAT,
// 		false,
// 		0,
// 		unsafe.Offsetof(card.CardVertices.Coords),
// 	)
// 
// 	gl.EnableVertexAttribArray(1)
// 	gl.VertexAttribPointerWithOffset(
// 		1,
// 		2,
// 		gl.FLOAT,
// 		false,
// 		0,
// 		unsafe.Offsetof(card.CardVertices.TexCoords),
// 	)
// 
// 	var indexBufferId uint32
// 	gl.GenBuffers(1, &indexBufferId)
// 	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBufferId)
// 
// 	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
// 		int(unsafe.Sizeof(card.CardVertexIndices)),
// 		gl.Ptr(card.CardVertexIndices[:]),
// 		gl.STATIC_DRAW,
// 	)
// 
// 	// Setup shaders
// 
// 	s.ProgramId = gl.CreateProgram()
// 
// 	s.attachShader(
// 		gl.VERTEX_SHADER,
// 		"resources/shaders/cards/vert.glsl",
// 	)
// 
// 	s.attachShader(
// 		gl.FRAGMENT_SHADER,
// 		"resources/shaders/cards/frag.glsl",
// 	)
// 
// 	gl.LinkProgram(s.ProgramId)
// 	gl.ValidateProgram(s.ProgramId)
// 	gl.UseProgram(s.ProgramId)
// 
// 	s.ProjectionUniformLocation = gl.GetUniformLocation(
// 		s.ProgramId,
// 		gl.Str("u_Projection\000"),
// 	)
// 
// 	gl.UniformMatrix4fv(
// 		s.ProjectionUniformLocation,
// 		1,
// 		false,
// 		&s.World.Projection[0],
// 	)
// 
// 	s.CameraUniformLocation = gl.GetUniformLocation(
// 		s.ProgramId,
// 		gl.Str("u_Camera\000"),
// 	)
// 
// 	s.TransformUniformLocation = gl.GetUniformLocation(
// 		s.ProgramId,
// 		gl.Str("u_Transform\000"),
// 	)
// 
// 	// Filled on demand
// 	s.CardTextures = make(map[string]uint32)
// 
// 	texSpriteLocation := gl.GetUniformLocation(
// 		s.ProgramId,
// 		gl.Str("texSprite\000"),
// 	)
// 
// 	gl.Uniform1i(texSpriteLocation, 0)
// 
// 	// Back side
// 	// TODO: Support eggs.
// 
// 	gl.ActiveTexture(digigl.CardBackTextureUnit)
// 
// 	var backTextureId uint32
// 	gl.GenTextures(1, &backTextureId)
// 	gl.BindTexture(gl.TEXTURE_2D, backTextureId)
// 
// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
// 	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
// 
// 	texture := s.loadTexture("back1")
// 	width := int32(texture.Bounds().Max.X)
// 	height := int32(texture.Bounds().Max.Y)
// 
// 	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0,
// 		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(texture.Pix))
// 	gl.GenerateMipmap(gl.TEXTURE_2D)
// 
// 	texBackLocation := gl.GetUniformLocation(
// 		s.ProgramId,
// 		gl.Str("texBack\000"),
// 	)
// 
// 	gl.Uniform1i(texBackLocation, 1)
// }
// 
// func (s *CardRendererSystem) attachShader(
// 	xtype uint32,
// 	path resources.ResPath,
// ) {
// 	shader, err := resources.ReadShader(path)
// 
// 	if err != nil {
// 		log.Fatalln("Could not load vertex shader for cards", err)
// 	}
// 
// 	shaderId := gl.CreateShader(xtype)
// 	shaderCompat, free := gl.Strs(shader)
// 
// 	shaderLength := int32(len(shader))
// 	gl.ShaderSource(shaderId, 1, shaderCompat, &shaderLength)
// 	gl.CompileShader(shaderId)
// 	free()
// 
// 	gl.AttachShader(s.ProgramId, shaderId)
// 	gl.DeleteShader(shaderId)
// }
// 
// func loadTexture(code string) resources.Texture {
// 	data, err := resources.MakeTexture(
// 		resources.ResPath(fmt.Sprintf("resources/textures/cards/%s.jpg", code)))
// 
// 	if err != nil {
// 		log.Fatalln("Could not load texture for card", code, "", err)
// 	}
// 
// 	return data
// }
// 
// // Allocates a new texture if the card hasn't been seen before or
// // simply returns the texture ID if it has.
// func (s *CardRendererSystem) getTextureId(code string) uint32 {
// 	if textureId, ok := s.CardTextures[code]; ok {
// 		return textureId
// 	}
// 
// 	texture := loadTexture(code)
// 	textureId := texture.GLGen(digigl.SpriteTextureUnit)
// 
// 	s.CardTextures[code] = textureId
// 	return textureId
// }
// 
// func (s *CardRendererSystem) Render(
// 	ctx engine.RenderContext,
// 	node engine.Node,
// ) {
// 	transform :=
// 		(*node.GetComponent(&components.TransformComponentKind)).(components.Transform)
// 	render :=
// 		(*node.GetComponent(&components.CardRendererComponentKind)).(components.CardRendererData)
// 
// 	gl.BindBuffer(gl.ARRAY_BUFFER, s.VertexBufferId)
// 	gl.BindVertexArray(s.VertexArrayId)
// 	defer gl.BindVertexArray(0)
// 
// 	gl.UseProgram(s.ProgramId)
// 
// 	gl.ActiveTexture(digigl.SpriteTextureUnit.GL())
// 	gl.BindTexture(gl.TEXTURE_2D, render.TextureId)
// 
// 	transformM := transform.ToMatrix()
// 	camera := ctx.World.CameraMatrix()
// 
// 	gl.UniformMatrix4fv(
// 		s.CameraUniformLocation,
// 		1,
// 		false,
// 		&camera[0],
// 	)
// 
// 	gl.UniformMatrix4fv(
// 		s.TransformUniformLocation,
// 		1,
// 		false,
// 		&transformM[0],
// 	)
// 
// 	gl.DrawElements(
// 		gl.TRIANGLES,
// 		int32(len(card.CardVertexIndices)),
// 		gl.UNSIGNED_INT,
// 		nil,
// 	)
// }
