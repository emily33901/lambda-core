package common

import (
	entity2 "github.com/emily33901/lambda-core/core/entity"
	"github.com/emily33901/lambda-core/game/entity"
)

// PropDynamic
type PropDynamic struct {
	entity2.Base
	entity.PropBase
}

// New
func (entity *PropDynamic) New() entity2.IEntity {
	return &PropDynamic{}
}

// Classname
func (entity PropDynamic) Classname() string {
	return "prop_dynamic"
}
