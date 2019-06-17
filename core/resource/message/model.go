package message

import (
	"github.com/emily33901/lambda-core/core/event"
	"github.com/emily33901/lambda-core/core/model"
)

const (
	// TypeModelLoaded
	TypeModelLoaded = event.MessageType("ModelLoaded")
	// TypeModelUnloaded
	TypeModelUnloaded = event.MessageType("ModelUnloaded")
)

// PropLoaded
type PropLoaded struct {
	event.Message
	// Resource
	Resource *model.Model
}

// Type
func (message *PropLoaded) Type() event.MessageType {
	return TypeModelLoaded
}

// PropUnloaded
type PropUnloaded struct {
	event.Message
	// Resource
	Resource *model.Model
}

// Type
func (message *PropUnloaded) Type() event.MessageType {
	return TypeModelUnloaded
}

// LoadedModel
func LoadedModel(mod *model.Model) event.IMessage {
	return &PropLoaded{
		Resource: mod,
	}
}

// UnloadedModel
func UnloadedModel(mod *model.Model) event.IMessage {
	return &PropUnloaded{
		Resource: mod,
	}
}
