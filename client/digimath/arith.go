package digimath

import "math"

func IsZero(f float32) bool {
	return math.Abs(float64(f)) < 1e-6
}

func Inf() float32 {
	return float32(math.Inf(1))
}
