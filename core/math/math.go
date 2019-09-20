package math

import (
	"github.com/go-gl/mathgl/mgl32"
)

func IntersectSegmentTriangle(segmentOrigin, segmentVec, vert0, vert1, vert2 mgl32.Vec3) (mgl32.Vec3, bool) {
	const epsilon = 0.0000001

	edge1 := vert1.Sub(vert0)
	edge2 := vert2.Sub(vert1)

	h := segmentVec.Cross(edge2)
	a := edge1.Dot(h)

	if a > -epsilon && a < epsilon {
		// parallel
		return mgl32.Vec3{0, 0, 0}, false
	}

	f := 1 / a
	s := segmentOrigin.Sub(vert0)
	u := f * s.Dot(h)

	if u < 0 || u > 1 {
		return mgl32.Vec3{0, 0, 0}, false
	}

	q := s.Cross(edge1)
	v := f * segmentVec.Dot(q)

	if v < 0.0 || u+v > 1.0 {
		return mgl32.Vec3{0, 0, 0}, false
	}

	t := f * edge2.Dot(q)

	if t > epsilon {
		// we hit
		return segmentOrigin.Add(segmentVec.Mul(t)), true
	}

	return mgl32.Vec3{0, 0, 0}, false
}

func VecAsSlice(v mgl32.Vec3) []float32 {
	return []float32{v.X(), v.Y(), v.Z()}
}

func Vec3sAsSlice(v []mgl32.Vec3) (result []float32) {
	result = make([]float32, len(v)*3)
	for i, x := range v {
		result[i*3] = x[0]
		result[i*3+1] = x[1]
		result[i*3+2] = x[2]
	}
	return result
}

func Vec4sAsSlice(v []mgl32.Vec4) (result []float32) {
	result = make([]float32, len(v)*4)
	for i, x := range v {
		result[i*3] = x[0]
		result[i*3+1] = x[1]
		result[i*3+2] = x[2]
		result[i*3+3] = x[3]
	}
	return result
}

func Vec2sAsSlice(v []mgl32.Vec2) (result []float32) {
	result = make([]float32, len(v)*2)
	for i, x := range v {
		result[i*2] = x[0]
		result[i*2+1] = x[1]
	}
	return result
}
