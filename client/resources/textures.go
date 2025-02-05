package resources

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/CrimsonSarah/cto/client/digigl"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Only holds image data.
type Texture struct {
	Data *image.RGBA
}

// Is associated with a texture unit.
type BoundTexture struct {
	Unit digigl.TextureUnit
	GLTextureId uint32
}

// Only JPEGs supported for now
func MakeTexture(path ResPath) (Texture, error) {
	filePath := GameFilePath(path)
	file, err := os.Open(filePath)

	if err != nil {
		return Texture{}, err
	}

	cardImage, err := jpeg.Decode(file)

	if err != nil {
		return Texture{}, err
	}

	data := image.NewRGBA(cardImage.Bounds())
	draw.Draw(
		data,
		data.Bounds(),
		cardImage,
		cardImage.Bounds().Min,
		draw.Src,
	)

	tex := Texture{
		Data: data,
	}

	return tex, nil
}

func (t Texture) Bounds() image.Rectangle {
	return t.Data.Bounds()
}

func MakeBoundTexture(
	tex Texture,
	unit digigl.TextureUnit,
) BoundTexture {
	btex := BoundTexture{
		Unit: unit,
	}

	gl.ActiveTexture(unit.GL())

	gl.GenTextures(1, &btex.GLTextureId)
	gl.BindTexture(gl.TEXTURE_2D, btex.GLTextureId)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	data := tex.Data
	width := int32(data.Bounds().Max.X)
	height := int32(data.Bounds().Max.Y)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(data.Pix))

	gl.GenerateMipmap(gl.TEXTURE_2D)

	return btex
}

func (bt *BoundTexture) Bind() {
	gl.ActiveTexture(bt.Unit.GL())
	gl.BindTexture(gl.TEXTURE_2D, bt.GLTextureId)
}
