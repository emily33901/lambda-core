package loader

import (
	entity3 "github.com/emily33901/lambda-core/core/entity"
)

// entityClassMapper provides a reflection-like construct for creating
// new entity objects of a known Classname.
// The idea behind this was to remove the need to slow, difficult to read
// reflection. Instead, it is up to defined entity types to provide a means
// to create a new instance of its own type; this class being used to provide
// a gateway to manage that mapping.
// Note: this class is somewhat memory costly, as a single unmodified instance for every
// mapped type is required for storage. Templated functions would probably solve this
// problem better if they existed, and the plan was to avoid actual reflection
// where possible.
type entityClassMapper struct {
	entityMap map[string]entity3.IEntity
}

// find creates a new Entity of the specified
// Classname.
func (classMap *entityClassMapper) find(classname string) entity3.IEntity {
	if classMap.entityMap[classname] != nil {
		return classMap.entityMap[classname].New()
	}
	return nil
}

var classMap entityClassMapper

// RegisterClass adds any type that implements a classname to
// a saved mapping. From then on, new instances of that classname
// can be created from just knowing the classname at runtime.
func RegisterClass(entity entity3.IClassname) {
	if classMap.entityMap == nil {
		classMap.entityMap = map[string]entity3.IEntity{}
	}

	classMap.entityMap[entity.Classname()] = entity.(entity3.IEntity)
}

// New creates a new Entity of the specified
// Classname.
func New(classname string) entity3.IEntity {
	return classMap.find(classname)
}
