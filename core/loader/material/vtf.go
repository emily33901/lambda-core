package material

import (
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda-Core/core/resource"
	"github.com/galaco/Lambda-Core/core/texture"
	"github.com/galaco/vtf"
	"strings"
)

// LoadSingleTexture
func LoadSingleTexture(filePath string, fs filesystem.IFileSystem) texture.ITexture {
	filePath = filesystem.NormalisePath(filePath)
	if !strings.HasSuffix(filePath, filesystem.ExtensionVtf) {
		filePath = filePath + filesystem.ExtensionVtf
	}
	if resource.Manager().Texture(filesystem.BasePathMaterial+filePath) != nil {
		return resource.Manager().Texture(filesystem.BasePathMaterial + filePath).(texture.ITexture)
	}
	if filePath == "" {
		return resource.Manager().Texture(resource.Manager().ErrorTextureName()).(texture.ITexture)
	}
	mat, err := readVtf(filesystem.BasePathMaterial+filePath, fs)
	if err != nil {
		logger.Warn("Failed to load texture: %s. Reason: %s", filesystem.BasePathMaterial+filePath, err)
		return resource.Manager().Texture(resource.Manager().ErrorTextureName()).(texture.ITexture)
	}
	return mat
}

// readVtf
func readVtf(path string, fs filesystem.IFileSystem) (texture.ITexture, error) {
	ResourceManager := resource.Manager()
	stream, err := fs.GetFile(path)
	if err != nil {
		return nil, err
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	read, err := vtf.ReadFromStream(stream)
	if err != nil {
		return nil, err
	}
	// Store filesystem containing raw data in memory
	ResourceManager.AddTexture(
		texture.NewTexture2D(
			path,
			read,
			int(read.GetHeader().Width),
			int(read.GetHeader().Height)))

	// Finally generate the gpu buffer for the material
	return ResourceManager.Texture(path).(texture.ITexture), nil
}
