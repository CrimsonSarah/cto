package digimath

import (
	"fmt"
	"math"
)

// Everything column-major. 32-bit floats to play well with the GPU.

// Vec2

type Vec2 [2]float32

func MakeVec2(x, y float32) Vec2 {
	return Vec2([2]float32{x, y})
}

func (v Vec2) X() float32 {
	return v[0]
}

func (v Vec2) U() float32 {
	return v[0]
}

func (v Vec2) Y() float32 {
	return v[1]
}

func (v Vec2) V() float32 {
	return v[1]
}

func (v Vec2) Scale(amount float32) Vec2 {
	return MakeVec2(v.X()*amount, v.Y()*amount)
}

// Vec3

type Vec3 [3]float32

var Vec3Zero = Vec3([3]float32{0, 0, 0})

func MakeVec3(x, y, z float32) Vec3 {
	return Vec3([3]float32{x, y, z})
}

func (v Vec3) X() float32 {
	return v[0]
}

func (v Vec3) Y() float32 {
	return v[1]
}

func (v Vec3) Z() float32 {
	return v[2]
}

func (v Vec3) Scale(amount float32) Vec3 {
	return MakeVec3(v.X()*amount, v.Y()*amount, v.Z()*amount)
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return MakeVec3(v1.X()+v2.X(), v1.Y()+v2.Y(), v1.Z()+v2.Z())
}

// Vec4

type Vec4 [4]float32

func MakeVec4(x, y, z, w float32) Vec4 {
	return Vec4([4]float32{x, y, z, w})
}

func (v Vec4) X() float32 {
	return v[0]
}

func (v Vec4) Y() float32 {
	return v[1]
}

func (v Vec4) Z() float32 {
	return v[2]
}

func (v Vec4) W() float32 {
	return v[3]
}

func (v Vec4) Scale(amount float32) Vec4 {
	return MakeVec4(v.X()*amount, v.Y()*amount, v.Z()*amount, v.W()*amount)
}

// Matrix33

type Matrix33 [9]float32

func MakeMatrix33(
	x11, x12, x13,
	x21, x22, x23,
	x31, x32, x33 float32,
) Matrix33 {
	return Matrix33([9]float32{
		x11, x21, x31,
		x12, x22, x32,
		x13, x23, x33,
	})
}

func Matrix33Id() Matrix33 {
	return MakeMatrix33(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)
}

// Matrix44 ❤️

type Matrix44 [16]float32

func MakeMatrix44(
	x11, x12, x13, x14,
	x21, x22, x23, x24,
	x31, x32, x33, x34,
	x41, x42, x43, x44 float32,
) Matrix44 {
	return Matrix44([16]float32{
		x11, x21, x31, x41,
		x12, x22, x32, x42,
		x13, x23, x33, x43,
		x14, x24, x34, x44,
	})
}

// 0 indexed!!
// TODO: ^ Don't
func (m Matrix44) Entry(i, j uintptr) float32 {
	return m[j*4+i]
}

func (m Matrix44) Scale(s float32) Matrix44 {
	return MakeMatrix44(
		m.Entry(0, 0)*s, m.Entry(0, 1)*s, m.Entry(0, 2)*s, m.Entry(0, 3)*s,
		m.Entry(1, 0)*s, m.Entry(1, 1)*s, m.Entry(1, 2)*s, m.Entry(1, 3)*s,
		m.Entry(2, 0)*s, m.Entry(2, 1)*s, m.Entry(2, 2)*s, m.Entry(2, 3)*s,
		m.Entry(3, 0)*s, m.Entry(3, 1)*s, m.Entry(3, 2)*s, m.Entry(3, 3)*s,
	)
}

func (m Matrix44) Mul(other Matrix44) Matrix44 {
	me := func(i, j uintptr) float32 {
		return m.Entry(i, 0)*other.Entry(0, j) +
			m.Entry(i, 1)*other.Entry(1, j) +
			m.Entry(i, 2)*other.Entry(2, j) +
			m.Entry(i, 3)*other.Entry(3, j)
	}

	return MakeMatrix44(
		me(0, 0), me(0, 1), me(0, 2), me(0, 3),
		me(1, 0), me(1, 1), me(1, 2), me(1, 3),
		me(2, 0), me(2, 1), me(2, 2), me(2, 3),
		me(3, 0), me(3, 1), me(3, 2), me(3, 3),
	)
}

func (m Matrix44) MulV(v Vec4) Vec4 {
	me := func(i uintptr) float32 {
		return m.Entry(i, 0)*v.X() +
			m.Entry(i, 1)*v.Y() +
			m.Entry(i, 2)*v.Z() +
			m.Entry(i, 3)*v.W()
	}

	return MakeVec4(
		me(0),
		me(1),
		me(2),
		me(3),
	)
}

func (m Matrix44) Transpose() Matrix44 {
	return MakeMatrix44(
		m.Entry(0, 0), m.Entry(1, 0), m.Entry(2, 0), m.Entry(3, 0),
		m.Entry(0, 1), m.Entry(1, 1), m.Entry(2, 1), m.Entry(3, 1),
		m.Entry(0, 2), m.Entry(1, 2), m.Entry(2, 2), m.Entry(3, 2),
		m.Entry(0, 3), m.Entry(1, 3), m.Entry(2, 3), m.Entry(3, 3),
	)
}

func Matrix44Id() Matrix44 {
	return MakeMatrix44(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func Matrix44Scale(amount float32) Matrix44 {
	return MakeMatrix44(
		amount, 0, 0, 0,
		0, amount, 0, 0,
		0, 0, amount, 0,
		0, 0, 0, amount,
	)
}

func Matrix44Translate(amount Vec3) Matrix44 {
	return MakeMatrix44(
		1, 0, 0, amount.X(),
		0, 1, 0, amount.Y(),
		0, 0, 1, amount.Z(),
		0, 0, 0, 1,
	)
}

func Matrix44RotateZ(amount float32) Matrix44 {
	cos := float32(math.Cos(float64(amount)))
	sin := float32(math.Sin(float64(amount)))

	return MakeMatrix44(
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func Matrix44RotateY(amount float32) Matrix44 {
	cos := float32(math.Cos(float64(amount)))
	sin := float32(math.Sin(float64(amount)))

	return MakeMatrix44(
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1,
	)
}

func Matrix44RotateX(amount float32) Matrix44 {
	cos := float32(math.Cos(float64(amount)))
	sin := float32(math.Sin(float64(amount)))

	return MakeMatrix44(
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1,
	)
}

// For debugging.
func (m *Matrix44) Format() string {
	return fmt.Sprintf(
		""+
			"[ %.2f, %.2f, %.2f, %.2f ]\n"+
			"[ %.2f, %.2f, %.2f, %.2f ]\n"+
			"[ %.2f, %.2f, %.2f, %.2f ]\n"+
			"[ %.2f, %.2f, %.2f, %.2f ]",
		m.Entry(0, 0), m.Entry(0, 1), m.Entry(0, 2), m.Entry(0, 3),
		m.Entry(1, 0), m.Entry(1, 1), m.Entry(1, 2), m.Entry(1, 3),
		m.Entry(2, 0), m.Entry(2, 1), m.Entry(2, 2), m.Entry(2, 3),
		m.Entry(3, 0), m.Entry(3, 1), m.Entry(3, 2), m.Entry(3, 3),
	)
}
