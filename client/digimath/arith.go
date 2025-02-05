package digimath

import "math"

func IsZero(f float32) bool {
	return math.Abs(float64(f)) < 1e-6
}

// Get rid of things like `9.536743e-07`.
func FixZero(f float32) float32 {
	if IsZero(f) {
		return 0
	} else {
		return f
	}
}

func Inf() float32 {
	return float32(math.Inf(1))
}
