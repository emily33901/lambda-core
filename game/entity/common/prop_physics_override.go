package common

import (
	"github.com/emily33901/lambda-core/core/entity"
	entity2 "github.com/emily33901/lambda-core/game/entity"
)

// PropPhysicsOverride
type PropPhysicsOverride struct {
	entity.Base
	entity2.PropBase
}

//New
func (entity *PropPhysicsOverride) New() entity.IEntity {
	return &PropPhysicsOverride{}
}

// Classname
func (entity PropPhysicsOverride) Classname() string {
	return "prop_physics_override"
}
