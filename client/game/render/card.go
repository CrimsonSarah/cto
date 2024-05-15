package render

import (
	"fmt"
	"image"
	"log"
	"unsafe"

	"github.com/CrimsonSarah/cto/client/digigl"
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/card"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/resources"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type CardRenderData struct {
	TextureId uint32
}

type RenderableCard struct {
	*world.Placed[card.Card]
	Render CardRenderData
}

type Vertex struct {
	x float32
	y float32
	z float32

	texX float32
	texY float32
}

var vertexBuffer = [4]Vertex{
	{x: -0.358, y: -0.5, z: 0, texX: 0, texY: 1},
	{x: 0.358, y: -0.5, z: 0, texX: 1, texY: 1},
	{x: 0.358, y: 0.5, z: 0, texX: 1, texY: 0},
	{x: -0.358, y: 0.5, z: 0, texX: 0, texY: 0},
}

var indexBuffer = [6]uint32{
	0, 1, 2, 2, 3, 0,
}

type CardRenderer struct {
	Projection *digimath.Matrix44

	VertexArrayId uint32
	ProgramId     uint32

	ProjectionUniformLocation           int32
	TransformScaleUniformLocation       int32
	TransformRotationUniformLocation    int32
	TransformTranslationUniformLocation int32

	CardTextures map[string]uint32
}

func (r *CardRenderer) Init(projection *digimath.Matrix44) {
	r.Projection = projection

	// Setup vertex and index buffers
	log.Println("Initialiazing cards")

	gl.GenVertexArrays(1, &r.VertexArrayId)
	gl.BindVertexArray(r.VertexArrayId)

	var vertexBufferId uint32
	gl.GenBuffers(1, &vertexBufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferId)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		int(unsafe.Sizeof(vertexBuffer)),
		gl.Ptr(vertexBuffer[:]),
		gl.STATIC_DRAW,
	)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(
		0,
		3,
		gl.FLOAT,
		false,
		int32(unsafe.Sizeof(Vertex{})),
		unsafe.Offsetof(Vertex{}.x),
	)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(
		1,
		2,
		gl.FLOAT,
		false,
		int32(unsafe.Sizeof(Vertex{})),
		unsafe.Offsetof(Vertex{}.texX),
	)

	var indexBufferId uint32
	gl.GenBuffers(1, &indexBufferId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBufferId)

	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		int(unsafe.Sizeof(indexBuffer)),
		gl.Ptr(indexBuffer[:]),
		gl.STATIC_DRAW,
	)

	// Setup shaders

	r.ProgramId = gl.CreateProgram()

	r.attachShader(
		gl.VERTEX_SHADER,
		"resources/shaders/cards/vert.glsl",
	)

	r.attachShader(
		gl.FRAGMENT_SHADER,
		"resources/shaders/cards/frag.glsl",
	)

	gl.LinkProgram(r.ProgramId)
	gl.ValidateProgram(r.ProgramId)
	gl.UseProgram(r.ProgramId)

	r.ProjectionUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_Projection\000"),
	)

	gl.UniformMatrix4fv(r.ProjectionUniformLocation, 1, false, &projection[0])

	// TODO: Read about UBOs
	r.TransformScaleUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_TransformScale\000"),
	)

	r.TransformRotationUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_TransformRotation\000"),
	)

	r.TransformTranslationUniformLocation = gl.GetUniformLocation(
		r.ProgramId,
		gl.Str("u_TransformTranslation\000"),
	)

	// Filled on demand
	r.CardTextures = make(map[string]uint32)
}

func (r *CardRenderer) Configure(projection *digimath.Matrix44) {
	gl.UseProgram(r.ProgramId)
	gl.UniformMatrix4fv(r.ProjectionUniformLocation, 1, false, &projection[0])
}

func (r *CardRenderer) attachShader(
	xtype uint32,
	path resources.ResPath,
) {
	shader, err := resources.ReadShader(path)

	if err != nil {
		log.Fatalln("Could not load vertex shader for cards", err)
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

func (r *CardRenderer) loadTexture(code string) *image.RGBA {
	data, err := resources.ReadTexture(
		resources.ResPath(fmt.Sprintf("resources/textures/cards/%s.jpg", code)))

	if err != nil {
		log.Fatalln("Could not load texture for card", code, "", err)
	}

	return data
}

// Allocates a new texture if the card hasn't been seen before or
// simply returns the texture ID if it has
func (r *CardRenderer) getTextureId(code string) uint32 {
	if textureId, ok := r.CardTextures[code]; ok {
		return textureId
	}

	gl.ActiveTexture(digigl.SpriteTextureUnit)

	var textureId uint32
	gl.GenTextures(1, &textureId)
	gl.BindTexture(gl.TEXTURE_2D, textureId)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	texture := r.loadTexture(code)
	width := int32(texture.Bounds().Max.X)
	height := int32(texture.Bounds().Max.Y)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(texture.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	r.CardTextures[code] = textureId
	return textureId
}

func (r *CardRenderer) MakeRenderableCard(card *world.Placed[card.Card]) RenderableCard {
	textureId := r.getTextureId(card.Obj.Code)

	result := RenderableCard{
		Placed: card,
		Render: CardRenderData{
			TextureId: textureId,
		},
	}

	return result
}

func (r *CardRenderer) RenderCard(c *RenderableCard) {
	gl.BindVertexArray(r.VertexArrayId)
	gl.UseProgram(r.ProgramId)

	gl.ActiveTexture(digigl.SpriteTextureUnit)
	gl.BindTexture(gl.TEXTURE_2D, c.Render.TextureId)

	gl.UniformMatrix4fv(
		r.TransformScaleUniformLocation,
		1,
		false,
		&c.Transform.ScaleMatrix[0],
	)

	gl.UniformMatrix4fv(
		r.TransformRotationUniformLocation,
		1,
		false,
		&c.Transform.RotationMatrix[0],
	)

	gl.UniformMatrix4fv(
		r.TransformTranslationUniformLocation,
		1,
		false,
		&c.Transform.TranslationMatrix[0],
	)

	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(indexBuffer)),
		gl.UNSIGNED_INT,
		nil,
	)
}
