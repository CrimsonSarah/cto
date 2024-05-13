package world

import "github.com/CrimsonSarah/cto/client/digimath"

// In degrees.
const PROJECTION_FOV_X float32 = 90

const PROJECTION_FAR float32 = 200
const PROJECTION_NEAR float32 = 0.1

func GetProjection(areaWidth, areaHeight float32) digimath.Matrix44 {
	// halfWidth := areaWidth / 2
	// halfHeight := areaHeight / 2
	halfWidth := float32(1)
	halfHeight := float32(1)

	return digimath.MakeMatrix44(
		1/halfWidth, 0, 0, 0,
		0, 1/halfHeight, 0, 0,
		0, 0, (-2/PROJECTION_FAR - PROJECTION_NEAR), -(PROJECTION_FAR+PROJECTION_NEAR)/(PROJECTION_FAR-PROJECTION_NEAR),
		0, 0, 0, 1,
	)
}

// func GetProjection(areaWidth, areaHeight float32) digimath.Matrix44 {
// 	aspectRatio := max(areaWidth, areaHeight) / min(areaWidth, areaHeight)
//
// 	// http://www.songho.ca/opengl/gl_projectionmatrix.html#hfov
// 	tangent := digimath.DegreesToRadians(PROJECTION_FOV_X)
// 	right := PROJECTION_NEAR * tangent
// 	top := right / aspectRatio
//
// 	return digimath.MakeMatrix44(
// 		PROJECTION_NEAR/right, 0, 0, 0,
// 		0, PROJECTION_NEAR/top, 0, 0,
// 		0, 0, -(PROJECTION_FAR+PROJECTION_NEAR)/(PROJECTION_FAR-PROJECTION_NEAR), -1,
// 		0, 0, -(2*PROJECTION_FAR*PROJECTION_NEAR)/(PROJECTION_FAR-PROJECTION_NEAR), 0,
// 	).Transpose()
// }
