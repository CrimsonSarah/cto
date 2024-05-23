package main

import (
	"fmt"
	"math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/game/world"
)

// For testing matrix operations because maths is hard.
// Maybe delete this at some point.

func main() {
	// TestProj()
	// TestInverses()
	// TestTransform()
	// TestProjInverses()
	// TestTransformInverses()
}

// Kind of emulates the vertex buffer
func vert(
	projection, transform digimath.Matrix44,
	pos digimath.Vec3,
) digimath.Vec4 {
	pos4 := digimath.MakeVec4(
		pos.X(),
		pos.Y(),
		pos.Z(),
		1,
	)

	projected := projection.Mul(transform).MulV(pos4)
	depth := projected.W()

	projected = projected.Scale(1 / depth)
	projected.SetW(depth)

	return projected
}

func TestProj() {
	fmt.Println("")
	fmt.Println("# Test1")

	v1 := digimath.MakeVec3(2, 3, -4)
	v2 := digimath.MakeVec3(-2, -3, -5)

	transform := digimath.Matrix44RotateX(0).Mul(
		digimath.Matrix44Id(),
	)

	projection := world.GetProjection(1000, 1000)
	noitcejorp := world.GetNoitcejorp(projection)
	noitcejorp2 := projection.Inverse()

	vv1 := vert(projection, transform, v1)
	vv2 := vert(projection, transform, v2)

	fmt.Printf("Projection\n%s\n", projection.Format())
	fmt.Printf("Projection^-1\n%s\n", noitcejorp.Format())
	fmt.Printf("Projection^-1 #2\n%s\n", noitcejorp2.Format())
	fmt.Printf("Transform\n%s\n", transform.Format())

	fmt.Println("")
	fmt.Printf("V1\n%v\n", v1)
	fmt.Printf("Result\n%v\n", vv1)

	fmt.Println("")
	fmt.Printf("V2\n%v\n", v2)
	fmt.Printf("Result\n%v\n", vv2)

	vv1Depth := vv1.W()
	vv1_2 := vv1.Scale(vv1Depth)
	vv1_2.SetW(vv1Depth)
	vv1_2 = noitcejorp.MulV(vv1_2)

	vv2Depth := vv2.W()
	vv2_2 := vv2.Scale(vv2Depth)
	vv2_2.SetW(vv2Depth)
	vv2_2 = noitcejorp.MulV(vv2_2)

	fmt.Println("")
	fmt.Printf("V1 %v\n", vv1_2)

	fmt.Println("")
	fmt.Printf("V2 %v\n", vv2_2)
}

func TestInverses() {
	fmt.Println("")
	fmt.Println("# Test2")

	mat1 := digimath.MakeMatrix44(
		2, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 2, 0,
		0, 0, 0, 2,
	)

	tam1 := digimath.MakeMatrix44(
		0.5, 0, 0, 0,
		0, 0.5, 0, 0,
		0, 0, 0.5, 0,
		0, 0, 0, 0.5,
	)

	mat2 := digimath.MakeMatrix44(
		1, 0, 0, 2,
		0, 1, 0, 2,
		0, 0, 1, 2,
		0, 0, 0, 1,
	)

	tam2 := digimath.MakeMatrix44(
		1, 0, 0, -2,
		0, 1, 0, -2,
		0, 0, 1, -2,
		0, 0, 0, 1,
	)

	mat21 := mat2.Mul(mat1)
	tam21 := tam2.Mul(tam1)

	id1 := tam1.Mul(mat1)
	id2 := tam2.Mul(mat2)
	id := tam21.Mul(mat21)

	fmt.Printf("Mat2 * Mat1\n%s\n", mat21.Format())
	fmt.Printf("Mat1 * Tam1\n%s\n", id1.Format())
	fmt.Printf("Tam2 * Tam1\n%s\n", tam21.Format())
	fmt.Printf("Mat2 * Tam2\n%s\n", id2.Format())
	fmt.Printf("(Tam2 * Tam1) * (Mat2 * Mat1) \n%s\n", id.Format())
}

func TestTransform() {
	fmt.Println("")
	fmt.Println("# Test3")

	s1 := float32(1)
	p1 := digimath.MakeVec3(1, 2, 3)
	r1 := digimath.MakeVec3(math.Pi/2, math.Pi/2, 0)

	t1 := world.MakeTransform()
	t1.ScaleFactor = s1
	t1.Position = p1
	t1.Rotation = r1

	mat1 := t1.ToMatrix()

	mat2 := digimath.Matrix44Id()
	mat2 = digimath.Matrix44RotateZ(r1.Z()).Mul(mat2)
	mat2 = digimath.Matrix44RotateY(r1.Y()).Mul(mat2)
	mat2 = digimath.Matrix44RotateX(r1.X()).Mul(mat2)
	mat2 = digimath.Matrix44Translate(p1).Mul(mat2)
	mat2 = digimath.Matrix44Scale(s1).Mul(mat2)

	fmt.Printf("Mat1\n%s\n", mat1.Format())
	fmt.Printf("Mat2\n%s\n", mat2.Format())

	v1 := digimath.MakeVec4(1, 1, 1, 0)

	fmt.Printf("Mat1 * V1\n%v\n", mat1.MulV(v1))
	fmt.Printf("Mat2 * V1\n%v\n", mat2.MulV(v1))
}

func TestProjInverses() {
	fmt.Println("")
	fmt.Println("# Test4")

	projection := world.GetProjection(600, 1000)
	noitcejorp := world.GetNoitcejorp(projection)

	fmt.Printf("Projection\n%s\n", projection.Format())
	fmt.Printf("Projection^-1\n%s\n", noitcejorp.Format())

	id1 := noitcejorp.Mul(projection)
	id2 := projection.Mul(noitcejorp)

	fmt.Println("")
	fmt.Printf("ID 1\n%s\n", id1.Format())
	fmt.Printf("ID 2\n%s\n", id2.Format())
}

func TestTransformInverses() {
	fmt.Println("")
	fmt.Println("# Test5")

	s1 := float32(1)
	p1 := digimath.MakeVec3(1, 2, 3)
	r1 := digimath.MakeVec3(math.Pi/6, math.Pi/6, math.Pi/4)

	t1 := world.MakeTransform()
	t1.ScaleFactor = s1
	t1.Position = p1
	t1.Rotation = r1

	mat1 := t1.ToMatrix()
	tam1_1 := mat1.Inverse()
	tam1_2 := mat1.Inverse0001()

	fmt.Printf("Mat1\n%s\n", mat1.Format())
	fmt.Printf("Mat1^-1\n%s\n", tam1_1.Format())
	fmt.Printf("Mat1^-1\n%s\n", tam1_2.Format())

	id1 := tam1_1.Mul(mat1)
	id2 := tam1_2.Mul(mat1)

	fmt.Println("")
	fmt.Printf("ID\n%s\n", id1.Format())
	fmt.Printf("ID\n%s\n", id2.Format())
}
