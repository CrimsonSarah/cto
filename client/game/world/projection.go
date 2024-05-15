package world

import (
	"github.com/CrimsonSarah/cto/client/digimath"
)

// In degrees.
const PROJECTION_FOV_Y float32 = 90

const PROJECTION_FAR float32 = 20
const PROJECTION_NEAR float32 = 1

func GetProjection(areaWidth, areaHeight float32) digimath.Matrix44 {
	aspectRatio := max(areaWidth, areaHeight) / min(areaWidth, areaHeight)

	// http://www.songho.ca/opengl/gl_projectionmatrix.html#vfov
	tangent := digimath.DegreesToRadians(PROJECTION_FOV_Y / 2)
	top := PROJECTION_NEAR * tangent
	right := top * aspectRatio

	projection := digimath.MakeMatrix44(
		(PROJECTION_NEAR / right), 0, 0, 0,
		0, (PROJECTION_NEAR / top), 0, 0,
		0, 0, -1, (-2 * PROJECTION_NEAR),
		0, 0, -1, 0,
	)

	return projection
}
