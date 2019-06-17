package common

import (
	entity2 "github.com/emily33901/lambda-core/core/entity"
	"github.com/emily33901/lambda-core/game/entity"
)

// PropPhysics
type PropPhysics struct {
	entity2.Base
	entity.PropBase
}

// New
func (entity *PropPhysics) New() entity2.IEntity {
	return &PropPhysics{}
}

// Classname
func (entity PropPhysics) Classname() string {
	return "prop_physics"
}
