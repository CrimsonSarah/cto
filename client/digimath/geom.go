package digimath

// Ray

type Line struct {
	Origin    Vec3
	Direction Vec3
}

func MakeLine(origin, direction Vec3) Line {
	return Line{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Line) NormalizedDirection() Line {
	return MakeLine(r.Origin, r.Direction.Normalized())
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

// Intersections

func IntersectLinePlane(line Line, plane Plane) (bool, Vec3) {
	planedirection := plane.Dot(line.Direction.AsDirection())

	if IsZero(planedirection) {
		return false, Vec3Zero
	}

	planeorigin := plane.Dot(line.Origin.AsPoint())

	return true, line.Origin.Sub(
		line.Direction.Scale(planeorigin / planedirection),
	)
}
