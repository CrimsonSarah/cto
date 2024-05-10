package render

import (
	"fmt"
	"image"
	"log"
	"unsafe"

	"github.com/CrimsonSarah/cto/client/card"
	"github.com/CrimsonSarah/cto/client/resources"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type RenderData struct {
	TextureId uint32
}

type RenderableCard struct {
	*card.PlacedCard
	Render RenderData
}

func MakeRenderableCard(card *card.PlacedCard) RenderableCard {
	result := RenderableCard{
		PlacedCard: card,
	}

	allocateTexture(&result)
	return result
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

var vertexArrayId uint32
var programId uint32
var transformUniformLocation int32

func CardInit() {
	// Setup vertex and index buffers

	gl.GenVertexArrays(1, &vertexArrayId)
	gl.BindVertexArray(vertexArrayId)

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

	programId = gl.CreateProgram()

	attachShader(
		programId,
		gl.VERTEX_SHADER,
		"data/shaders/cards/vert.glsl",
	)

	attachShader(programId,
		gl.FRAGMENT_SHADER,
		"data/shaders/cards/frag.glsl",
	)

	gl.LinkProgram(programId)
	gl.ValidateProgram(programId)
	gl.UseProgram(programId)

	transformUniformLocation = gl.GetUniformLocation(
		programId,
		gl.Str("u_Transform\000"),
	)
}

func attachShader(
	programId uint32,
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

	gl.AttachShader(programId, shaderId)
	gl.DeleteShader(shaderId)
}

func getTexture(c *RenderableCard) *image.RGBA {
	data, err := resources.ReadTexture(
		resources.ResPath(fmt.Sprintf("data/textures/cards/%s.jpg", c.Card.Id)))

	if err != nil {
		log.Fatalln("Could not load texture for card", c.Card.Id, "", err)
	}

	return data
}

func allocateTexture(c *RenderableCard) {
	gl.ActiveTexture(SpriteTextureUnit)
	gl.GenTextures(1, &c.Render.TextureId)
	gl.BindTexture(gl.TEXTURE_2D, c.Render.TextureId)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	texture := getTexture(c)
	width := int32(texture.Bounds().Max.X)
	height := int32(texture.Bounds().Max.Y)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(texture.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
}

func RenderCard(c *RenderableCard) {
	gl.BindVertexArray(vertexArrayId)
	gl.UseProgram(programId)

	gl.ActiveTexture(SpriteTextureUnit)
	gl.BindTexture(gl.TEXTURE_2D, c.Render.TextureId)

	gl.UniformMatrix4fv(transformUniformLocation, 1, false, &c.Transform[0])

	gl.DrawElements(
		gl.TRIANGLES,
		int32(len(indexBuffer)),
		gl.UNSIGNED_INT,
		nil,
	)
}
