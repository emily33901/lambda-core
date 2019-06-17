package scene

import (
	"github.com/emily33901/lambda-core/core/model"
)

// IScene
type IScene interface {
	// Bsp
	Bsp() *model.Bsp
	// StaticProps
	StaticProps() []model.StaticProp
}
