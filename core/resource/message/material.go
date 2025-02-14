package message

import (
	"github.com/emily33901/lambda-core/core/event"
	"github.com/emily33901/lambda-core/core/material"
)

const (
	// TypeMaterialLoaded
	TypeMaterialLoaded = event.MessageType("MaterialLoaded")
	// TypeMaterialUnloaded
	TypeMaterialUnloaded = event.MessageType("MaterialUnloaded")
)

// MaterialLoaded
type MaterialLoaded struct {
	event.Message
	// Resource
	Resource material.IMaterial
}

// Type
func (message *MaterialLoaded) Type() event.MessageType {
	return TypeMaterialLoaded
}

// MaterialUnloaded
type MaterialUnloaded struct {
	event.Message
	// Resource
	Resource material.IMaterial
}

// Type
func (message *MaterialUnloaded) Type() event.MessageType {
	return TypeMaterialUnloaded
}

// LoadedMaterial
func LoadedMaterial(mat material.IMaterial) event.IMessage {
	return &MaterialLoaded{
		Resource: mat,
	}
}

// UnloadedMaterial
func UnloadedMaterial(mat material.IMaterial) event.IMessage {
	return &MaterialUnloaded{
		Resource: mat,
	}
}
