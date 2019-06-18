package entity

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const cameraSpeed = float64(320)
const sensitivity = float32(0.03)

var minVerticalRotation = mgl32.DegToRad(90)
var maxVerticalRotation = mgl32.DegToRad(270)

const (
	OrthoX = 0
	OrthoY = 1
	OrthoZ = 2
)

var (
	OrthoDirections = [...]string{
		"Front Z/Y",
		"Top X/Z",
		"Size X/Y",
	}
)

// Camera
type Camera struct {
	*Base
	fov              float32
	aspectRatio      float32
	up               mgl32.Vec3
	right            mgl32.Vec3
	direction        mgl32.Vec3
	worldUp          mgl32.Vec3
	ortho            bool
	orthoOrientation int
	orthoZoom        float32
}

// Forwards
func (camera *Camera) Forwards(dt float64) {
	if !camera.ortho {
		camera.Transform().Position = camera.Transform().Position.Add(camera.direction.Mul(float32(cameraSpeed * dt)))
	} else {
		camera.Transform().Position = camera.Transform().Position.Add(camera.up.Mul(float32(-cameraSpeed * dt)))
	}
}

// Backwards
func (camera *Camera) Backwards(dt float64) {
	if !camera.ortho {
		camera.Transform().Position = camera.Transform().Position.Sub(camera.direction.Mul(float32(cameraSpeed * dt)))
	} else {
		camera.Transform().Position = camera.Transform().Position.Sub(camera.up.Mul(float32(-cameraSpeed * dt)))
	}
}

// Left
func (camera *Camera) Left(dt float64) {
	camera.Transform().Position = camera.Transform().Position.Sub(camera.right.Mul(float32(cameraSpeed * dt)))
}

// Right
func (camera *Camera) Right(dt float64) {
	camera.Transform().Position = camera.Transform().Position.Add(camera.right.Mul(float32(cameraSpeed * dt)))
}

// Move
func (camera *Camera) Move(dt mgl32.Vec3) {
	camera.Transform().Position = camera.Transform().Position.Add(dt)
}

// Rotate
func (camera *Camera) Rotate(x, y, z float32) {
	camera.Transform().Rotation[0] -= float32(x * sensitivity)
	camera.Transform().Rotation[1] -= float32(y * sensitivity)
	camera.Transform().Rotation[2] -= float32(z * sensitivity)

	// Lock vertical rotation
	if camera.Transform().Rotation[2] > maxVerticalRotation {
		camera.Transform().Rotation[2] = maxVerticalRotation
	}
	if camera.Transform().Rotation[2] < minVerticalRotation {
		camera.Transform().Rotation[2] = minVerticalRotation
	}
}

func (camera *Camera) setRotation(x, y, z float32) {
	camera.Transform().Rotation[0] = float32(x)
	camera.Transform().Rotation[1] = float32(y)
	camera.Transform().Rotation[2] = float32(z)
}

// SetAspect
func (camera *Camera) SetAspect(newAspect float32) {
	camera.aspectRatio = newAspect
}

func (camera *Camera) Zoom(dt float32) {
	if !camera.ortho {
		camera.fov += dt

		if camera.fov > 71 {
			camera.fov = 71
		} else if camera.fov < 69.3 {
			camera.fov = 69.3
		}
	} else {
		camera.orthoZoom += dt * 100
	}

}

func (camera *Camera) Ortho() bool {
	return camera.ortho
}

func (camera *Camera) SetOrtho(enabled bool) {
	camera.ortho = enabled
}

func (camera *Camera) OrthoDirection() int {
	return camera.orthoOrientation
}
func (camera *Camera) SetOrthoDirection(new int) {
	camera.orthoOrientation = new
}

// Update updates the camera position
func (camera *Camera) Update(dt float64) {
	camera.updateVectors()
}

// updateVectors Updates the camera directional properties with any changes
func (camera *Camera) updateVectors() {
	rot := camera.Transform().Rotation

	// Calculate the new Front vector
	camera.direction = mgl32.Vec3{
		float32(math.Cos(float64(rot[2])) * math.Sin(float64(rot[0]))),
		float32(math.Cos(float64(rot[2])) * math.Cos(float64(rot[0]))),
		float32(math.Sin(float64(rot[2]))),
	}
	// Also re-calculate the right and up vector
	camera.right = mgl32.Vec3{
		float32(math.Sin(float64(rot[0]) - math.Pi/2)),
		float32(math.Cos(float64(rot[0]) - math.Pi/2)),
		0,
	}
	camera.up = camera.right.Cross(camera.direction)
}

// ScreenToWorldMatrix returns the matrix that can transform a screen point into its
// world coordinates
func (camera *Camera) ScreenToWorld(viewPort, point mgl32.Vec2) mgl32.Vec3 {
	r, err := mgl32.UnProject(point.Vec3(0.5), camera.ModelMatrix().Mul4(camera.ViewMatrix()), camera.ProjectionMatrix(), 0, 0, int(viewPort[0]), int(viewPort[1]))

	if err != nil {
		panic(err)
	}

	return r

	// mvp := camera.ModelMatrix().Mul4(camera.ProjectionMatrix()).Mul4(camera.ViewMatrix()).Inv()

	// clipSpace := mvp.Mul(1 / mvp[3])

	// // Convert the point to ndc's
	// w := viewPort.X()
	// h := viewPort.Y()

	// point[0] -= 0.5 * w
	// point[1] -= 0.5 * h

	// point[0] = (point[0] / w) * 2
	// point[1] = (point[1] / w) * 2

	// point4 := point.Vec4(0, 0)
	// point4[3] = 0

	// return mvp.Mul4x1(point4)
}

// ModelMatrix returns identity matrix (camera model is our position!)
func (camera *Camera) ModelMatrix() mgl32.Mat4 {
	return mgl32.Ident4()
}

// ViewMatrix calculates the cameras View matrix
func (camera *Camera) ViewMatrix() mgl32.Mat4 {
	if camera.ortho {
		// @TODO We probably shouldnt be doing this every time...
		// And we defo should not be doing it here!
		switch camera.orthoOrientation {
		case OrthoX:
			camera.setRotation(0, mgl32.DegToRad(-90), 0)
		case OrthoY:
			camera.setRotation(0, 0, 0)
		case OrthoZ:
			camera.setRotation(mgl32.DegToRad(90), 0, 0)
		}
	}

	return mgl32.LookAtV(
		camera.Transform().Position,
		camera.Transform().Position.Add(camera.direction),
		camera.up)
}

// ProjectionMatrix calculates projection matrix.
func (camera *Camera) ProjectionMatrix() mgl32.Mat4 {
	if camera.ortho {
		base := 50 * camera.orthoZoom
		return mgl32.Ortho(0, base*camera.aspectRatio, base, 0, -99999, 99999)
	}

	return mgl32.Perspective(camera.fov, camera.aspectRatio, 0.1, 16384)
}

func (camera *Camera) Fov() float32 {
	return camera.fov
}

// NewCamera returns a new camera
// fov should be provided in radians
func NewCamera(fov float32, aspectRatio float32) *Camera {
	return &Camera{
		Base:        &Base{},
		fov:         fov,
		aspectRatio: aspectRatio,
		up:          mgl32.Vec3{0, 1, 0},
		worldUp:     mgl32.Vec3{0, 1, 0},
		direction:   mgl32.Vec3{0, 0, -1},
		orthoZoom:   70,
	}
}
