package digigl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

var SpriteTextureUnit uint32
var CardBackTextureUnit uint32

// Should be called once OpenGL is initialized
func TextureInit() {
	var offset = 0

	SpriteTextureUnit = uint32(gl.TEXTURE0 + offset)
	offset += 1

	CardBackTextureUnit = uint32(gl.TEXTURE0 + offset)
	// offset += 1
}
