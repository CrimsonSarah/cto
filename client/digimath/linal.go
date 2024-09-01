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

func Vec3From4(vec Vec4) Vec3 {
	return MakeVec3(vec.X(), vec.Y(), vec.Z())
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

func (v Vec3) MagnitudeSquared() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec3) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.MagnitudeSquared())))
}

func (v Vec3) Normalized() Vec3 {
	return v.Scale(1 / v.Magnitude())
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return MakeVec3(v1.X()+v2.X(), v1.Y()+v2.Y(), v1.Z()+v2.Z())
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return MakeVec3(v1.X()-v2.X(), v1.Y()-v2.Y(), v1.Z()-v2.Z())
}

func (v1 Vec3) Dot(v2 Vec3) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return MakeVec3(
		v1.Y()*v2.Z()-v1.Z()*v2.Y(),
		v1.Z()*v2.X()-v1.X()*v2.Z(),
		v1.X()*v2.Y()-v1.Y()*v2.X(),
	)
}

func (v1 Vec3) AsPoint() Vec4 {
	return MakeVec4(v1.X(), v1.Y(), v1.Z(), 1)
}

func (v1 Vec3) AsDirection() Vec4 {
	return MakeVec4(v1.X(), v1.Y(), v1.Z(), 0)
}

// Vec4

type Vec4 [4]float32

func MakeVec4(x, y, z, w float32) Vec4 {
	return Vec4([4]float32{x, y, z, w})
}

func (v Vec4) X() float32 {
	return v[0]
}
func (v *Vec4) SetX(x float32) {
	v[0] = x
}

func (v Vec4) Y() float32 {
	return v[1]
}
func (v *Vec4) SetY(y float32) {
	v[1] = y
}

func (v Vec4) Z() float32 {
	return v[2]
}
func (v *Vec4) SetZ(z float32) {
	v[2] = z
}

func (v Vec4) W() float32 {
	return v[3]
}
func (v *Vec4) SetW(w float32) {
	v[3] = w
}

func (v Vec4) Scale(amount float32) Vec4 {
	return MakeVec4(v.X()*amount, v.Y()*amount, v.Z()*amount, v.W()*amount)
}

func (v1 Vec4) Add(v2 Vec4) Vec4 {
	return MakeVec4(v1.X()+v2.X(), v1.Y()+v2.Y(), v1.Z()+v2.Z(), v1.W()+v2.W())
}

func (v1 Vec4) Sub(v2 Vec4) Vec4 {
	return MakeVec4(v1.X()-v2.X(), v1.Y()-v2.Y(), v1.Z()-v2.Z(), v1.W()-v2.W())
}

func (v1 Vec4) Dot(v2 Vec4) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z() + v1.W()*v2.W()
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

// 1 indexed. `i` and `j` should never be less than 1.
// There are no bound checks of course.
func (m Matrix44) Entry(i, j uintptr) float32 {
	return m[(j-1)*4+(i-1)]
}

func (m Matrix44) Column(j uintptr) Vec4 {
	// Can't do `reinterpret_cast`-type stuff here, I don't think
	return MakeVec4(
		m.Entry(1, j),
		m.Entry(2, j),
		m.Entry(3, j),
		m.Entry(4, j),
	)
}

func (m Matrix44) Row(i uintptr) Vec4 {
	return MakeVec4(
		m.Entry(i, 1),
		m.Entry(i, 2),
		m.Entry(i, 3),
		m.Entry(i, 4),
	)
}

func (m Matrix44) Scale(s float32) Matrix44 {
	return MakeMatrix44(
		m.Entry(1, 1)*s, m.Entry(1, 2)*s, m.Entry(1, 3)*s, m.Entry(1, 4)*s,
		m.Entry(2, 1)*s, m.Entry(2, 2)*s, m.Entry(2, 3)*s, m.Entry(2, 4)*s,
		m.Entry(3, 1)*s, m.Entry(3, 2)*s, m.Entry(3, 3)*s, m.Entry(3, 4)*s,
		m.Entry(4, 1)*s, m.Entry(4, 2)*s, m.Entry(4, 3)*s, m.Entry(4, 4)*s,
	)
}

func (m Matrix44) Mul(other Matrix44) Matrix44 {
	me := func(i, j uintptr) float32 {
		return m.Entry(i, 1)*other.Entry(1, j) +
			m.Entry(i, 2)*other.Entry(2, j) +
			m.Entry(i, 3)*other.Entry(3, j) +
			m.Entry(i, 4)*other.Entry(4, j)
	}

	return MakeMatrix44(
		me(1, 1), me(1, 2), me(1, 3), me(1, 4),
		me(2, 1), me(2, 2), me(2, 3), me(2, 4),
		me(3, 1), me(3, 2), me(3, 3), me(3, 4),
		me(4, 1), me(4, 2), me(4, 3), me(4, 4),
	)
}

func (m Matrix44) MulV(v Vec4) Vec4 {
	me := func(i uintptr) float32 {
		return m.Entry(i, 1)*v.X() +
			m.Entry(i, 2)*v.Y() +
			m.Entry(i, 3)*v.Z() +
			m.Entry(i, 4)*v.W()
	}

	return MakeVec4(
		me(1),
		me(2),
		me(3),
		me(4),
	)
}

func (m Matrix44) Transpose() Matrix44 {
	return MakeMatrix44(
		m.Entry(1, 1), m.Entry(2, 1), m.Entry(3, 1), m.Entry(4, 1),
		m.Entry(1, 2), m.Entry(2, 2), m.Entry(3, 2), m.Entry(4, 2),
		m.Entry(1, 3), m.Entry(2, 3), m.Entry(3, 3), m.Entry(4, 3),
		m.Entry(1, 4), m.Entry(2, 4), m.Entry(3, 4), m.Entry(4, 4),
	)
}

// Generic and expensive.
func (m Matrix44) Inverse() Matrix44 {
	a := Vec3From4(m.Column(1))
	b := Vec3From4(m.Column(2))
	c := Vec3From4(m.Column(3))
	d := Vec3From4(m.Column(4))

	x := m.Entry(4, 1)
	y := m.Entry(4, 2)
	z := m.Entry(4, 3)
	w := m.Entry(4, 4)

	s := a.Cross(b)
	t := c.Cross(d)
	u := a.Scale(y).Sub(b.Scale(x))
	v := c.Scale(w).Sub(d.Scale(z))

	invDet := 1 / (s.Dot(v) + t.Dot(u))
	s = s.Scale(invDet)
	t = t.Scale(invDet)
	u = u.Scale(invDet)
	v = v.Scale(invDet)

	r1 := b.Cross(v).Add(t.Scale(y))
	r2 := v.Cross(a).Sub(t.Scale(x))
	r3 := d.Cross(u).Add(s.Scale(w))
	r4 := u.Cross(c).Sub(s.Scale(z))

	return MakeMatrix44(
		r1.X(), r1.Y(), r1.Z(), -b.Dot(t),
		r2.X(), r2.Y(), r2.Z(), a.Dot(t),
		r3.X(), r3.Y(), r3.Z(), -d.Dot(s),
		r4.X(), r4.Y(), r4.Z(), c.Dot(s),
	)
}

// Slightly optimized for the case where the last row is known to be
// [0 0 0 1].
// No checks of course.
func (m Matrix44) Inverse0001() Matrix44 {
	a := Vec3From4(m.Column(1))
	b := Vec3From4(m.Column(2))
	c := Vec3From4(m.Column(3))
	d := Vec3From4(m.Column(4))

	s := a.Cross(b)
	t := c.Cross(d)
	u := Vec3Zero
	v := c

	invDet := 1 / (s.Dot(v) + t.Dot(u))
	s = s.Scale(invDet)
	t = t.Scale(invDet)
	u = u.Scale(invDet)
	v = v.Scale(invDet)

	r1 := b.Cross(v)
	r2 := v.Cross(a)
	r3 := s
	r4 := u.Cross(c)

	return MakeMatrix44(
		r1.X(), r1.Y(), r1.Z(), -b.Dot(t),
		r2.X(), r2.Y(), r2.Z(), a.Dot(t),
		r3.X(), r3.Y(), r3.Z(), -d.Dot(s),
		r4.X(), r4.Y(), r4.Z(), c.Dot(s),
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
		m.Entry(1, 1), m.Entry(1, 2), m.Entry(1, 3), m.Entry(1, 4),
		m.Entry(2, 1), m.Entry(2, 2), m.Entry(2, 3), m.Entry(2, 4),
		m.Entry(3, 1), m.Entry(3, 2), m.Entry(3, 3), m.Entry(3, 4),
		m.Entry(4, 1), m.Entry(4, 2), m.Entry(4, 3), m.Entry(4, 4),
	)
}
