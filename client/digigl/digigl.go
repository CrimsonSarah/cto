// Package for interacting with OpenGL across all possible GLAreas.
package digigl

import (
	"log"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func openGlDebug(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	_ int32,
	message string,
	_ unsafe.Pointer) {

	log.Printf("[OpenGL Debug Message] source: %d, type: %d, id: %d, severity: %d, message: %s\n",
		source, gltype, id, severity, message)
}

func DigiGLInit() error {
	if err := gl.Init(); err != nil {
		log.Fatalln("Could not initialize OpenGL", err)
		return err
	}

	gl.Enable(gl.DEBUG_OUTPUT)
	gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
	gl.Enable(gl.BLEND)

	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.DebugMessageCallback(openGlDebug, nil)
	gl.DebugMessageControl(gl.DONT_CARE, gl.DONT_CARE, gl.DONT_CARE, 0, nil, true)

	gl.ClearColor(0, 0, 0, 0)

	TextureInit()
	log.Println("DigiGL Initialized")
	return nil
}
