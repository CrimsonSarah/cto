package digigl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type TextureUnit uint32

var SpriteTextureUnit TextureUnit
var CardBackTextureUnit TextureUnit

// Should be called once OpenGL is initialized
func TextureInit() {
	var offset = 0

	SpriteTextureUnit = TextureUnit(gl.TEXTURE0 + offset)
	offset += 1

	CardBackTextureUnit = TextureUnit(gl.TEXTURE0 + offset)
	// offset += 1
}

func (u TextureUnit) GL() uint32 {
	return uint32(u)
}
