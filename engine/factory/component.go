package factory

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/interfaces"
)

// Attaches a component to an entity, and registers it with the engine
func NewComponent(component interfaces.IComponent, owner interfaces.IEntity) *interfaces.IComponent {
	component.SetHandle(core.NewHandle())
	GetObjectManager().AddComponent(component, owner)
	return &component
}
