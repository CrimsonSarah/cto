package camera

// Provides utilities for working with the Camera from other systems,
// using the fact that the Camera entity ID is stored globally.

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
	"github.com/CrimsonSarah/cto/client/game/generic/transform"
	"github.com/CrimsonSarah/cto/client/game/global"
)

// The position of the camera.
func GetCameraTransform(ctx engine.IEngine[global.GameGlobals]) transform.Transform {
	cameraEntityId := ctx.GetGlobals().Camera
	t, ok := ctx.
		GetComponents(transform.TransformComponentKind.Id).
		Get(int(cameraEntityId))

	if !ok {
		panic("Looked for camera transform and did not find it.")
	}

	return t.Data.(transform.Transform)
}

// The matrix you multiply an object's position by.
func GetCameraViewMatrix(ctx engine.IEngine[global.GameGlobals]) digimath.Matrix44 {
	transform := GetCameraTransform(ctx)
	return transform.ToMatrixInverse()
}

// Inverse of the view matrix.
func GetAremacViewMatrix(ctx engine.IEngine[global.GameGlobals]) digimath.Matrix44 {
	transform := GetCameraTransform(ctx)
	return transform.ToMatrix()
}

func GetCameraCamera(ctx engine.IEngine[global.GameGlobals]) CameraCamera {
	cameraEntityId := ctx.GetGlobals().Camera
	cc, ok := ctx.
		GetComponents(CameraCameraComponentKind.Id).
		Get(int(cameraEntityId))

	if !ok {
		panic("Looked for camera component and did not find it.")
	}

	return cc.Data.(CameraCamera)
}
