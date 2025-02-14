package message

import (
	"github.com/emily33901/lambda-core/core/event"
	"github.com/emily33901/lambda-core/core/texture"
)

const (
	// TypeTextureLoaded
	TypeTextureLoaded = event.MessageType("TextureLoaded")
	// TypeTextureUnloaded
	TypeTextureUnloaded = event.MessageType("TextureUnloaded")
)

// TextureLoaded
type TextureLoaded struct {
	event.Message
	// Resource
	Resource texture.ITexture
}

// Type
func (message *TextureLoaded) Type() event.MessageType {
	return TypeTextureLoaded
}

// TextureUnloaded
type TextureUnloaded struct {
	event.Message
	// Resource
	Resource texture.ITexture
}

// Type
func (message *TextureUnloaded) Type() event.MessageType {
	return TypeTextureUnloaded
}

// LoadedTexture
func LoadedTexture(tex texture.ITexture) event.IMessage {
	return &TextureLoaded{
		Resource: tex,
	}
}

// UnloadedTexture
func UnloadedTexture(tex texture.ITexture) event.IMessage {
	return &TextureUnloaded{
		Resource: tex,
	}
}
