package digimath

// Ray

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func MakeRay(origin, direction Vec3) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) NormalizedDirection() Ray {
	return MakeRay(r.Origin, r.Direction.Normalized())
}

// Plane

type Plane struct {
	// xyz are the plane's normal vector, w is the constant dot
	// product for points in it.
	Vec Vec4
}

func MakePlane(normal Vec3, dot float32) Plane {
	return Plane{
		Vec: MakeVec4(normal.X(), normal.Y(), normal.Z(), dot),
	}
}

func (p Plane) Normal() Vec3 {
	return Vec3From4(p.Vec)
}

func (p Plane) D() float32 {
	return p.Vec.W()
}

func (p Plane) NormalizedNormal() Plane {
	normal := p.Normal()
	magnitude := normal.Magnitude()
	return MakePlane(normal.Scale(1/magnitude), p.D()/magnitude)
}

func (p Plane) Dot(x Vec4) float32 {
	return p.Vec.Dot(x)
}

// Rect

type Rect struct {
	BottomLeft Vec3
	Side1      Vec3 // bottom left <-> bottom right
	Side2      Vec3 // bottom left <-> top left
}

func MakeRect(bottomLeft, side1, side2 Vec3) Rect {
	return Rect{
		BottomLeft: bottomLeft,
		Side1:      side1,
		Side2:      side2,
	}
}

// A rect centered on (0,0) parallel to the XY plane.
func MakeProfileRect(width, height float32) Rect {
	bottomLeft := MakeVec3(-width/2, -height/2, 0)
	side1 := MakeVec3(width, 0, 0)
	side2 := MakeVec3(0, height, 0)

	return Rect{
		BottomLeft: bottomLeft,
		Side1:      side1,
		Side2:      side2,
	}
}

func (r Rect) GetBottomLeft() Vec3 {
	return r.BottomLeft
}

func (r Rect) GetBottomRight() Vec3 {
	return r.BottomLeft.Add(r.Side1)
}

func (r Rect) GetTopLeft() Vec3 {
	return r.BottomLeft.Add(r.Side2)
}

func (r Rect) GetTopRight() Vec3 {
	return r.BottomLeft.Add(r.Side1.Add(r.Side2))
}

// TODO: Is this correct?
func (r Rect) ToPlane() Plane {
	normal := r.Side2.Cross(r.Side1).Normalized()
	d := -normal.Dot(r.BottomLeft)

	// log.Printf(
	// 	"Normal: %v. BottomLeft: %v. u: %v. v: %v. d: %v.",
	// 	normal.FixZero(),
	// 	r.BottomLeft.FixZero(),
	// 	r.Side1.FixZero(),
	// 	r.Side2.FixZero(),
	// 	FixZero(d),
	// )

	return Plane{
		Vec: MakeVec4(normal.X(), normal.Y(), normal.Z(), d),
	}
}

// Intersections

func IntersectRayPlane(Ray Ray, plane Plane) (bool, Vec3) {
	planedirection := plane.Dot(Ray.Direction.AsDirection())

	if IsZero(planedirection) {
		return false, Vec3Zero
	}

	planeorigin := plane.Dot(Ray.Origin.AsPoint())

	return true, Ray.Origin.Sub(
		Ray.Direction.Scale(planeorigin / planedirection),
	)
}

func IntersectRayRect(Ray Ray, rect Rect) (bool, Vec3) {
	plane := rect.ToPlane()
	planedirection := plane.Dot(Ray.Direction.AsDirection())

	if IsZero(planedirection) {
		return false, Vec3Zero
	}

	planeorigin := plane.Dot(Ray.Origin.AsPoint())
	point := Ray.Origin.Sub(
		Ray.Direction.Scale(planeorigin / planedirection),
	)

	// Translate so that the bottom left is on the origin.
	bottomLeft := rect.GetBottomLeft()
	point2 := point.Sub(bottomLeft)

	u := point2.Dot(rect.Side1)
	v := point2.Dot(rect.Side2)

	maxU := rect.Side1.MagnitudeSquared()
	maxV := rect.Side2.MagnitudeSquared()

	if 0 <= u && u <= maxU && 0 <= v && v <= maxV {
		return true, point
	} else {
		return false, point
	}
}
