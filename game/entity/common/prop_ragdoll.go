package common

import (
	entity2 "github.com/emily33901/lambda-core/core/entity"
	"github.com/emily33901/lambda-core/game/entity"
)

// PropRagdoll
type PropRagdoll struct {
	entity2.Base
	entity.PropBase
}

// New
func (entity *PropRagdoll) New() entity2.IEntity {
	return &PropRagdoll{}
}

// Classname
func (entity PropRagdoll) Classname() string {
	return "prop_ragdoll"
}
