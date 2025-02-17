package resource

import (
	"strings"
	"sync"

	"github.com/emily33901/lambda-core/core/event"
	"github.com/emily33901/lambda-core/core/filesystem"
	"github.com/emily33901/lambda-core/core/material"
	"github.com/emily33901/lambda-core/core/model"
	"github.com/emily33901/lambda-core/core/resource/message"
	"github.com/emily33901/lambda-core/core/texture"
)

// Very generic filesystem storage.
// If the struct came from a filesystem, it should be obtainable from here
type manager struct {
	errorModelName   string
	errorTextureName string

	materials         map[string]material.IMaterial
	materialReadMutex sync.RWMutex
	textures          map[string]texture.ITexture
	textureReadMutex  sync.RWMutex
	models            map[string]*model.Model
	modelReadMutex    sync.RWMutex
}

// Add a new material
func (m *manager) AddMaterial(file material.IMaterial) {
	if m.HasMaterial(file.FilePath()) {
		return
	}
	m.materialReadMutex.Lock()
	m.materials[strings.ToLower(file.FilePath())] = file
	m.materialReadMutex.Unlock()

	event.Manager().Dispatch(message.LoadedMaterial(file))
}

// Add a new material
func (m *manager) AddTexture(file texture.ITexture) {
	if m.HasTexture(file.FilePath()) {
		return
	}
	m.textureReadMutex.Lock()
	m.textures[strings.ToLower(file.FilePath())] = file
	m.textureReadMutex.Unlock()

	event.Manager().Dispatch(message.LoadedTexture(file))
}

// Add a new model
func (m *manager) AddModel(file *model.Model) {
	if m.HasModel(file.FilePath()) {
		return
	}
	m.modelReadMutex.Lock()
	m.models[strings.ToLower(file.FilePath())] = file
	m.modelReadMutex.Unlock()

	event.Manager().Dispatch(message.LoadedModel(file))
}

// Get Find a specific filesystem
func (m *manager) Material(filePath string) material.IMaterial {
	m.materialReadMutex.Lock()
	defer m.materialReadMutex.Unlock()

	return m.materials[strings.ToLower(filePath)]
}

func (m *manager) Texture(filePath string) texture.ITexture {
	m.textureReadMutex.Lock()
	defer m.textureReadMutex.Unlock()

	return m.textures[strings.ToLower(filePath)]
}

func (m *manager) Model(filePath string) *model.Model {
	m.modelReadMutex.Lock()
	defer m.modelReadMutex.Unlock()

	return m.models[strings.ToLower(filePath)]
}

func (m *manager) Materials() map[string]material.IMaterial {
	return m.materials
}

func (m *manager) Textures() map[string]texture.ITexture {
	return m.textures
}

func (m *manager) Models() map[string]*model.Model {
	return m.models
}

// ErrorModelName Get error model name
func (m *manager) ErrorModelName() string {
	return m.errorModelName
}

// SetErrorModelName Override the default error model.
// Useful for when HL2 assets are not available (they include the engine
// default model)
func (m *manager) SetErrorModelName(name string) {
	m.errorModelName = name
}

// ErrorTextureName Get error texture name
func (m *manager) ErrorTextureName() string {
	return m.errorTextureName
}

// SetErrorTextureName Override default error texture
func (m *manager) SetErrorTextureName(name string) {
	m.errorTextureName = name
}

// Has the specified file been loaded
func (m *manager) HasMaterial(filePath string) bool {
	m.materialReadMutex.Lock()
	defer m.materialReadMutex.Unlock()

	if m.materials[strings.ToLower(filePath)] != nil {
		return true
	}
	return false
}

func (m *manager) HasTexture(filePath string) bool {
	m.textureReadMutex.Lock()
	defer m.textureReadMutex.Unlock()

	if m.textures[strings.ToLower(filePath)] != nil {
		return true
	}
	return false
}

// Has the specified model been loaded
func (m *manager) HasModel(filePath string) bool {
	m.modelReadMutex.Lock()
	defer m.modelReadMutex.Unlock()

	if m.models[strings.ToLower(filePath)] != nil {
		return true
	}
	return false
}

func (m *manager) Empty() {
	for idx, val := range m.materials {
		event.Manager().Dispatch(message.UnloadedMaterial(val))
		delete(m.materials, idx)
	}
	for idx, val := range m.textures {
		event.Manager().Dispatch(message.UnloadedTexture(val))
		delete(m.textures, idx)
	}
	for idx, val := range m.models {
		event.Manager().Dispatch(message.UnloadedModel(val))
		delete(m.models, idx)
	}
}

var resourceManager manager

// Manager returns the static filemanager
func Manager() *manager {
	if resourceManager.materials == nil {
		resourceManager.errorModelName = filesystem.BasePathModels + "error.mdl"
		resourceManager.errorTextureName = filesystem.BasePathMaterial + "error" + filesystem.ExtensionVtf
		resourceManager.materials = make(map[string]material.IMaterial, 1024)
		resourceManager.models = make(map[string]*model.Model, 256)
		resourceManager.textures = make(map[string]texture.ITexture, 256)
	}

	return &resourceManager
}
