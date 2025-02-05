package camera

// This is the component where the camera stuff is, except for the
// Transform, which of course is a Transform component.

import (
	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/engine"
)

var CameraCameraComponentKind = engine.MakeComponentKind("CameraCamera")

type CameraCamera struct {
	WindowWidth  int
	WindowHeight int

	Projection digimath.Matrix44
	Noitcejorp digimath.Matrix44
}

func MakeCameraCamera(width, height int) CameraCamera {
	projection := GetProjection(
		float32(width),
		float32(height),
	)

	noitcejorp := GetNoitcejorp(projection)

	return CameraCamera{
		WindowWidth:  width,
		WindowHeight: height,
		Projection:   projection,
		Noitcejorp:   noitcejorp,
	}
}

func (c CameraCamera) GetKind() *engine.ComponentKind {
	return &CameraCameraComponentKind
}

func MakeCameraCameraComponent(width, height int) engine.Component {
	return engine.MakeComponent(MakeCameraCamera(width, height))
}

// Projection:

// In degrees.
const PROJECTION_FOV_Y float32 = 60

const PROJECTION_FAR float32 = 20
const PROJECTION_NEAR float32 = 1

// "Weak" perspective projection, according to Wikipedia.
func GetProjection(areaWidth, areaHeight float32) digimath.Matrix44 {
	aspectRatio := max(areaWidth, areaHeight) / min(areaWidth, areaHeight)

	// http://www.songho.ca/opengl/gl_projectionmatrix.html#ortho
	tangent := digimath.DegreesToRadians(PROJECTION_FOV_Y / 2)
	top := PROJECTION_NEAR * tangent
	right := top * aspectRatio

	e11 := 1 / right
	e22 := 1 / top
	e33 := -2 / (PROJECTION_FAR - PROJECTION_NEAR)
	e34 := -(PROJECTION_FAR + PROJECTION_NEAR) / (PROJECTION_FAR - PROJECTION_NEAR)

	projection := digimath.MakeMatrix44(
		e11, 0, 0, 0,
		0, e22, 0, 0,
		0, 0, e33, e34,
		0, 0, -1, 0,
	)

	return projection
}

// Projection inverse.
func GetNoitcejorp(projection digimath.Matrix44) digimath.Matrix44 {
	e11 := 1 / projection.Entry(1, 1)
	e22 := 1 / projection.Entry(2, 2)

	e34 := float32(-1)
	e43 := 1 / projection.Entry(3, 4)
	e44 := projection.Entry(3, 3) / projection.Entry(3, 4)

	noitcejorp := digimath.MakeMatrix44(
		e11, 0, 0, 0,
		0, e22, 0, 0,
		0, 0, 0, e34,
		0, 0, e43, e44,
	)

	return noitcejorp
}
