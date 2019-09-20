package mesh

import (
	"github.com/emily33901/lambda-core/core/material"
	"github.com/emily33901/lambda-core/core/texture"
	"github.com/go-gl/mathgl/mgl32"
)

// IMesh Generic Mesh interface
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IMesh interface {
	// AddVertex
	AddVertex(...mgl32.Vec3)
	// AddNormal
	AddNormal(...mgl32.Vec3)
	// AddUV
	AddUV(...mgl32.Vec2)
	// AddLightmapCoordinate
	AddLightmapCoordinate(...mgl32.Vec3)
	// GenerateTangents
	GenerateTangents()

	// Vertices
	Vertices() []mgl32.Vec3
	// Normals
	Normals() []mgl32.Vec3
	// UVs
	UVs() []mgl32.Vec2
	// Tangents
	Tangents() []mgl32.Vec4
	// Colors
	Colors() []float32
	ResetColors(colors ...float32)
	// LightmapCoordinates
	LightmapCoordinates() []mgl32.Vec3

	// Material
	Material() material.IMaterial
	// SetMaterial
	SetMaterial(material.IMaterial)
	// Lightmap
	Lightmap() texture.ITexture
	// SetLightmap
	SetLightmap(texture.ITexture)
}
