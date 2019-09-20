package mesh

import (
	"github.com/emily33901/lambda-core/core/material"
	"github.com/emily33901/lambda-core/core/mesh/util"
	"github.com/emily33901/lambda-core/core/texture"
	"github.com/go-gl/mathgl/mgl32"
)

// Mesh
type Mesh struct {
	vertices            []mgl32.Vec3
	normals             []mgl32.Vec3
	uvs                 []mgl32.Vec2
	tangents            []mgl32.Vec4
	lightmapCoordinates []mgl32.Vec3
	colors              []float32

	material material.IMaterial
	lightmap texture.ITexture
}

// AddVertex
func (mesh *Mesh) AddVertex(vertex ...mgl32.Vec3) {
	mesh.vertices = append(mesh.vertices, vertex...)
}

// AddNormal
func (mesh *Mesh) AddNormal(normal ...mgl32.Vec3) {
	mesh.normals = append(mesh.normals, normal...)
}

// AddUV
func (mesh *Mesh) AddUV(uv ...mgl32.Vec2) {
	mesh.uvs = append(mesh.uvs, uv...)
}

// AddTangent
func (mesh *Mesh) AddTangent(tangent ...mgl32.Vec4) {
	mesh.tangents = append(mesh.tangents, tangent...)
}

// AddLightmapCoordinate
func (mesh *Mesh) AddLightmapCoordinate(uv ...mgl32.Vec3) {
	mesh.lightmapCoordinates = append(mesh.lightmapCoordinates, uv...)
}

// Vertices
func (mesh *Mesh) Vertices() []mgl32.Vec3 {
	return mesh.vertices
}

// Normals
func (mesh *Mesh) Normals() []mgl32.Vec3 {
	return mesh.normals
}

// UVs
func (mesh *Mesh) UVs() []mgl32.Vec2 {
	return mesh.uvs
}

// Tangents
func (mesh *Mesh) Tangents() []mgl32.Vec4 {
	return mesh.tangents
}

// LightmapCoordinates
func (mesh *Mesh) LightmapCoordinates() []mgl32.Vec3 {
	// use standard uvs if there is no lightmap. Not ideal,
	// but there MUST be UVs, but they'll be ignored anyway if there is no
	// lightmap
	if len(mesh.lightmapCoordinates) == 0 {
		trans := func(a []mgl32.Vec2) (result []mgl32.Vec3) {
			result = make([]mgl32.Vec3, len(a))
			for i, x := range a {
				result[i] = x.Vec3(0.0)
			}

			return result
		}
		return trans(mesh.UVs())
	}
	return mesh.lightmapCoordinates
}

// Material
func (mesh *Mesh) Material() material.IMaterial {
	return mesh.material
}

// SetMaterial
func (mesh *Mesh) SetMaterial(mat material.IMaterial) {
	mesh.material = mat
}

// Lightmap
func (mesh *Mesh) Lightmap() texture.ITexture {
	return mesh.lightmap
}

//SetLightmap
func (mesh *Mesh) SetLightmap(mat texture.ITexture) {
	mesh.lightmap = mat
}

// GenerateTangents
func (mesh *Mesh) GenerateTangents() {
	mesh.tangents = util.GenerateTangents(mesh.Vertices(), mesh.Normals(), mesh.UVs())
}

func (mesh *Mesh) ResetColors(colors ...float32) {
	mesh.colors = colors
}

func (mesh *Mesh) AddColor(colors ...float32) {
	mesh.colors = append(mesh.colors, colors...)
}

func (mesh *Mesh) Colors() []float32 {
	return mesh.colors
}

// Higher level mesh features

// AddLine adds a line between 2 points
func (mesh *Mesh) AddLine(color []float32, a, b mgl32.Vec3) {
	mesh.AddVertex(a)
	mesh.AddVertex(b)
	mesh.AddVertex(b)
	mesh.AddNormal(mgl32.Vec3{0, 0, 0})
	mesh.AddNormal(mgl32.Vec3{0, 0, 0})
	mesh.AddNormal(mgl32.Vec3{0, 0, 0})
	mesh.AddUV(mgl32.Vec2{0, 0})
	mesh.AddUV(mgl32.Vec2{0, 0})
	mesh.AddUV(mgl32.Vec2{0, 0})
	mesh.AddColor(color...)
	mesh.AddColor(color...)
	mesh.AddColor(color...)
}

// NewMesh
func NewMesh() *Mesh {
	return &Mesh{}
}
